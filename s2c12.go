package main

import (
	"crypto/aes"
	"cryptopals/pkcs7"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
)

var GlobalKey []byte

func main() {
	keyString := flag.String("k", "0000000000000000", "AES key")
	myString := flag.String("s", "", "prefix string")
	flag.Parse()

	encodedString := `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK`

	unknownString, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("unknownString is %d bytes long\n", len(unknownString))

	encodedBytes := combineAndEncrypt([]byte(*myString), unknownString, []byte(*keyString))

	fmt.Printf("%d encoded bytes\n", len(encodedBytes))
}

func combineAndEncrypt(prefixString, unknownString, key []byte) []byte {
	cleartext := make([]byte, len(prefixString)+len(unknownString))
	var i, j int
	var b byte
	for i, b = range prefixString {
		cleartext[i] = b
	}
	for j, b = range unknownString {
		cleartext[i+j] = b
	}

	encodedBytes := aesECBencode(cleartext, key)

	return encodedBytes
}

func aesECBencode(cleartext []byte, key []byte) []byte {

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	paddedBytes := pkcs7.PadBlock(cleartext, aes.BlockSize)

	dst := make([]byte, len(paddedBytes))

	var i int
	for i = 0; i < len(paddedBytes); i += aes.BlockSize {
		block.Encrypt(dst[i:i+aes.BlockSize], paddedBytes[i:i+aes.BlockSize])
	}

	return dst
}
