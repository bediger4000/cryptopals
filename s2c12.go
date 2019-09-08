package main

import (
	"cryptopals/detect"
	"cryptopals/myaes"
	"cryptopals/xor"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"strings"
)

var GlobalKey []byte

func main() {
	keyString := flag.String("k", "YELLOW SUBMARINE", "AES key")
	flag.Parse()

	encodedString := `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK`
	// encodedString := `bm93IGlzIHRoZSB0aW1lIGZvciBhbGwgZ29vZCBtZW4K`
	// encodedString := `YWJjZGVmZ2hpamtsbW5vcAo=`
	// encodedString := `YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXo=`
	// encodedString := `YWJjZGVmZ2hpamtsb21ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1BRUlNUVVZXWFla`

	unknownString, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("unknownString is %d bytes long\n", len(unknownString))

	blockLength := detect.FindCipherBlockLength(myaes.ECBEncode, []byte(*keyString))
	if blockLength > -1 {
		fmt.Printf("Cipher block length %d\n", blockLength)
	} else {
		fmt.Printf("No detectable cipher block length\n")
		return
	}

	findUnencryptedBytes(myaes.ECBEncode, blockLength, []byte(*keyString), unknownString)
}

func findUnencryptedBytes(cipher func([]byte, []byte) []byte, blockLength int, key []byte, unknownString []byte) {
	var decipheredBytes []byte

	encryptedOriginal := cipher(unknownString, key)

	blockCount := len(cipher(unknownString, key)) / blockLength

	fmt.Printf("%d encrypted blocks\n", blockCount)

	for blockNo := 0; blockNo < blockCount; blockNo++ {

		for i := blockLength; i > 0; i-- {
			var prefixBuffer []byte
			for j := 0; j < i-1; j++ {
				prefixBuffer = append(prefixBuffer, 'A')
			}
			fmt.Printf("blockNo %d, prefix length %d, encrypting prefix %q\n", blockNo, i, string(prefixBuffer))

			encryptedBytes := combineAndEncrypt(cipher, prefixBuffer, unknownString, key)
			encryptedBytes = encryptedBytes[blockNo*blockLength : blockNo*blockLength+blockLength]
			fmt.Printf("Using encryptedBytes[%d:%d]\n", blockNo*blockLength, blockNo*blockLength+blockLength)

			var testBuffer []byte
			for j := 0; j < i-1; j++ {
				testBuffer = append(testBuffer, 'A')
			}
			for j := 0; j < len(decipheredBytes); j++ {
				testBuffer = append(testBuffer, decipheredBytes[j])
			}
			testBuffer = append(testBuffer, byte(0))
			fmt.Printf("Block number %d, block length %d\n", blockNo, blockLength)
			fmt.Printf("Test buffer, length(%d): %q", len(testBuffer), testBuffer)
			fmt.Printf("cutting out [%d:%d]\n", blockNo*blockLength, blockNo*blockLength+blockLength)
			testBuffer = testBuffer[blockNo*blockLength : blockNo*blockLength+blockLength]
			fmt.Printf("Test buffer, length(%d): %q\n", len(testBuffer), testBuffer)

			var candidate int
			for candidate = 0; candidate < 256; candidate++ {

				testBuffer[blockLength-1] = byte(candidate)
				candidateBlock := cipher(testBuffer, key)

				if xor.CompareBuffers(candidateBlock[0:blockLength], encryptedBytes) {
					decipheredBytes = append(decipheredBytes, byte(candidate))
					break
				}
			}
			if candidate == 256 {
				fmt.Printf("Block length %d, prefix length %d, block number %d\n", blockLength, i, blockNo)
				fmt.Printf("Did not find a byte that matched\n")
				fmt.Printf("%d bytes decrypted so far: %q\n", len(decipheredBytes), string(decipheredBytes))
				return
			}
			fmt.Printf("Decrypted so far: %q\n", string(decipheredBytes))
			encryptedTest := cipher(decipheredBytes, key)
			if xor.CompareBuffers(encryptedTest, encryptedOriginal) {
				fmt.Printf("Found it:\n%s\n", string(decipheredBytes))
				return
			}
		}
	}
}

func combineAndEncrypt(cipher func([]byte, []byte) []byte, prefixString, unknownString, key []byte) []byte {
	return cipher(append(prefixString, unknownString...), key)
}

func printBuffer(buffer []byte) string {
	sb := &strings.Builder{}
	divider := ""
	for _, b := range buffer {
		sb.WriteString(fmt.Sprintf("%s%3d", divider, b))
		divider = " "
	}
	return sb.String()
}
