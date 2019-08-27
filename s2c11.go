package main

import (
	"crypto/aes"
	"cryptopals/pkcs7"
	"cryptopals/xor"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

/*
An ECB/CBC detection oracle

Now that you have ECB and CBC working:

Write a function to generate a random AES key; that's just 16 random bytes.

Write a function that encrypts data under an unknown key --- that is, a
function that generates a random key and encrypts under it.

The function should look like:

encryption_oracle(your-input)
=> [MEANINGLESS JIBBER JABBER]

Under the hood, have the function append 5-10 bytes (count chosen randomly)
before the plaintext and 5-10 bytes after the plaintext.

Now, have the function choose to encrypt under ECB 1/2 the time, and under CBC
the other half (just use random IVs each time for CBC). Use rand(2) to decide
which to use.

Detect the block cipher mode the function is using each time. You should end up
with a piece of code that, pointed at a block box that might be encrypting ECB
or CBC, tells you which one is happening.

*/

func main() {

	fileName := flag.String("f", "", "name of file containing plaintext")
	cipherSelection := flag.Int("s", -1, "select CBC (1) or ECB (0) or random (default)")
	keySelection := flag.Int("k", -1, "select 16, 24 or 32 byte randomly-selected keys, default random choice")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	if *fileName == "" {
		*fileName = os.Args[1]
	}

	cleartext, err := ioutil.ReadFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	ciphertext, cipherMode := RandomizedCipher(cleartext, *cipherSelection, *keySelection)

	// Overkill
	if tryKeySizes(ciphertext, 4, 40) {
		fmt.Printf("Hamming distance says probably ECB\n")
	} else {
		fmt.Printf("Hamming distance says probably CBC\n")
	}

	if tryRepeatingBlocks(ciphertext) {
		fmt.Printf("Repeating blocks says probably ECB\n")
	} else {
		fmt.Printf("Repeating blocks says probably CBC\n")
	}
	fmt.Printf("Actually %s\n", cipherMode)
}

func tryRepeatingBlocks(ciphertext []byte) bool {

	limit := len(ciphertext) - aes.BlockSize

	for offset1 := 0; offset1 < limit; offset1 += aes.BlockSize {
		originalBlock := ciphertext[offset1 : offset1+aes.BlockSize]
		for offset2 := offset1 + aes.BlockSize; offset2 < limit; offset2 += aes.BlockSize {
			comparisonBlock := ciphertext[offset2 : offset2+aes.BlockSize]
			if compareBlocks(originalBlock, comparisonBlock) {
				return true
			}
		}
	}
	return false
}

var keylen = []int{16, 24, 32}

func RandomizedCipher(cleartext []byte, cipherSelection int, keySelection int) ([]byte, string) {

	modifiedCleartext := prependAndAppend(cleartext)

	var randomkey []byte

	if keySelection != 16 && keySelection != 24 && keySelection != 32 {
		randomkey = generateRandomKey(keylen[rand.Intn(3)])
	} else {
		randomkey = generateRandomKey(keySelection)
	}

	if cipherSelection == -1 {
		cipherSelection = rand.Intn(2)
	}

	var ciphertext []byte
	var cipherMode string
	switch cipherSelection {
	case 0:
		ciphertext = ECBEncode(modifiedCleartext, randomkey)
		cipherMode = "ECB"
	case 1:
		ciphertext = CBCEncode(modifiedCleartext, randomkey)
		cipherMode = "CBC"
	}

	return ciphertext, cipherMode
}

func prependAndAppend(cleartext []byte) []byte {
	Nbefore := 5 + rand.Intn(5)
	Nafter := 6 + rand.Intn(5)

	modifiedBuffer := make([]byte, Nbefore+len(cleartext)+Nafter)

	var i int
	for i = 0; i < Nbefore; i++ {
		modifiedBuffer[i] = byte(rand.Intn(256))
	}

	copy(modifiedBuffer[i:], cleartext)
	i += len(cleartext)

	for i = 0; i < Nafter; i++ {
		modifiedBuffer[i] = byte(rand.Intn(256))
	}

	return modifiedBuffer
}

func CBCEncode(bytes []byte, key []byte) []byte {

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	paddedBuffer := pkcs7.PadBlock(bytes, aes.BlockSize)

	dst := make([]byte, len(paddedBuffer))

	// Random initialization vector
	previousBlock := make([]byte, aes.BlockSize)
	for i := range previousBlock {
		previousBlock[i] = byte(rand.Intn(256))
	}

	for i := 0; i < len(bytes); i += aes.BlockSize {
		block.Encrypt(dst[i:i+aes.BlockSize], xor.Encode(paddedBuffer[i:i+aes.BlockSize], previousBlock))
		previousBlock = dst[i : i+aes.BlockSize]
	}

	return dst
}

func generateRandomKey(keylen int) []byte {
	randomkey := make([]byte, keylen)

	for i := 0; i < keylen; i++ {
		randomkey[i] = byte(rand.Intn(256))
	}

	return randomkey
}

func ECBEncode(bytes []byte, key []byte) []byte {

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	paddedBytes := pkcs7.PadBlock(bytes, aes.BlockSize)

	dst := make([]byte, len(paddedBytes))

	var i int
	for i = 0; i < len(bytes); i += aes.BlockSize {
		block.Encrypt(dst[i:i+aes.BlockSize], paddedBytes[i:i+aes.BlockSize])
	}

	return dst
}

func tryKeySizes(buffer []byte, minKeyLength, maxKeyLength int) bool {
	var sum float64
	var count int
	var hdAt2KeyLength float64
	for keyLength := minKeyLength; keyLength <= maxKeyLength; keyLength++ {
		sumBits, bufferCount := compareSubBuffers(keyLength, buffer)
		hammingDistance := float64(sumBits) / float64(bufferCount*keyLength)
		count++
		sum += hammingDistance
		if keyLength == 2*aes.BlockSize {
			hdAt2KeyLength = hammingDistance
		}
	}
	mean := sum / float64(count)
	if mean/hdAt2KeyLength >= 1.2 {
		return true
	}
	return false
}

func compareSubBuffers(keyLength int, buffer []byte) (int, int) {
	bufferLength := len(buffer)
	bufferComparisons := 0
	commonBitsCount := 0
	comparisonBuffer := buffer[0:keyLength]
	max := bufferLength - keyLength
	for j := keyLength; j < max; j += keyLength {
		xoredSubBuffer := xor.Encode(comparisonBuffer, buffer[j:j+keyLength])
		commonBitsCount += xor.CountBits(xoredSubBuffer)
		bufferComparisons++
	}
	return commonBitsCount, bufferComparisons
}

func compareBlocks(block1, block2 []byte) bool {
	for i, b := range block1 {
		if b != block2[i] {
			return false
		}
	}
	return true
}
