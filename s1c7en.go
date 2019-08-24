package main

import (
	"crypto/aes"
	"cryptopals/pkcs7"
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

	paddedBytes := pkcs7.PadBlock(bytes, aes.BlockSize)

	fmt.Fprintf(os.Stderr, "Allocating %d bytes for encrypted bytes\n", len(paddedBytes))
	dst := make([]byte, len(paddedBytes))

	var i int
	for i = 0; i < len(bytes); i += aes.BlockSize {
		block.Encrypt(dst[i:i+aes.BlockSize], paddedBytes[i:i+aes.BlockSize])
	}

	base64text := base64.StdEncoding.EncodeToString(dst)

	_, err = os.Stdout.Write([]byte(base64text))
	if err != nil {
		log.Fatal(err)
	}
}
