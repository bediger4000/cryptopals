package myaes

import (
	"crypto/aes"
	"cryptopals/pkcs7"
	"log"
)

func ECBEncode(cleartext []byte, key []byte) []byte {

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
