package main

/*
 The Base64-encoded content in this file has been encrypted via AES-128 in ECB
 mode under the key

"YELLOW SUBMARINE".

(case-sensitive, without the quotes; exactly 16 characters; I like "YELLOW
SUBMARINE" because it's exactly 16 bytes long, and now you do too).

Decrypt it. You know the key, after all.

Easiest way: use OpenSSL::Cipher and give it AES-128-ECB as the cipher.
*/

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	key := `YELLOW SUBMARINE`

	buffer, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := base64.StdEncoding.DecodeString(string(buffer))
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	if len(bytes)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	dst := make([]byte, len(bytes))

	for i := 0; i < len(bytes); i += aes.BlockSize {
		block.Decrypt(dst[i:i+aes.BlockSize], bytes[i:i+aes.BlockSize])
	}

	n, err := os.Stdout.Write(dst)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stderr, "Wrote %d bytes to stdout\n", n)
}
