package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	unknownString := `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK`

	ciphertext, err := base64.StdEncoding.DecodeString(unknownString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ciphertext is %d bytes long\n", len(ciphertext))
}
