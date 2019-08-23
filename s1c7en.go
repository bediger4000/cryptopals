package main

import (
	"crypto/aes"
	"cryptopals/blocks"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	key := `YELLOW SUBMARINE`

	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stderr, "%d bytes cleartext\n", len(bytes))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	if len(bytes)%aes.BlockSize != 0 {
		fmt.Fprintf(os.Stderr, "ciphertext length %d is not a multiple of the block size\n", len(bytes))
	}

	fmt.Fprintf(os.Stderr, "Allocating %d bytes for encrypted bytes\n", len(bytes)+(aes.BlockSize-len(bytes)%aes.BlockSize))
	dst := make([]byte, len(bytes)+(aes.BlockSize-len(bytes)%aes.BlockSize))

	var i int
	for i = 0; i < len(bytes); i += aes.BlockSize {
		block.Encrypt(dst[i:i+aes.BlockSize], bytes[i:i+aes.BlockSize])
	}

	if i != len(bytes) {
		bytesEncrypted := i - aes.BlockSize
		fmt.Fprintf(os.Stderr, "Encrypted %d bytes\n", bytesEncrypted)
		paddedBytes := blocks.Pkcs7Pad(bytes[bytesEncrypted:], aes.BlockSize)
		block.Encrypt(dst[bytesEncrypted:], paddedBytes[:])
	}

	base64text := base64.StdEncoding.EncodeToString(dst)

	_, err = os.Stdout.Write([]byte(base64text))
	if err != nil {
		log.Fatal(err)
	}
}
