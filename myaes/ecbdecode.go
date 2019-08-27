package myaes

import (
	"crypto/aes"
	"cryptopals/pkcs7"
	"log"
)

func ECBDecode(bytes []byte, key []byte) []byte {

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	bytesLen := len(bytes)
	if bytesLen%aes.BlockSize != 0 {
		// It's ok to panic here, bytes[] should have length
		// a multiple of block size.
		panic("ciphertext is not a multiple of the block size")
	}

	dst := make([]byte, bytesLen)

	for i := 0; i < bytesLen; i += aes.BlockSize {
		block.Decrypt(dst[i:i+aes.BlockSize], bytes[i:i+aes.BlockSize])
	}

	return pkcs7.TrimBuffer(dst)
}
