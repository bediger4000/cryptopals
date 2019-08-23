package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

/*

Implement PKCS#7 padding

A block cipher transforms a fixed-sized block (usually 8 or 16 bytes) of plaintext into ciphertext. But we almost never want to transform a single block; we encrypt irregularly-sized messages.

One way we account for irregularly-sized messages is by padding, creating a plaintext that is an even multiple of the blocksize. The most popular padding scheme is called PKCS#7.

So: pad any block to a specific block length, by appending the number of bytes of padding to the end of the block. For instance,

"YELLOW SUBMARINE"

... padded to 20 bytes would be:

"YELLOW SUBMARINE\x04\x04\x04\x04"
*/

func main() {
	block := os.Args[1]
	specificLength, err := strconv.Atoi(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}

	paddedBlock := pkcs7Pad([]byte(block), specificLength)

	fmt.Fprintf(os.Stderr, "length of padded block %d\n", len(paddedBlock))
	err = ioutil.WriteFile(os.Args[3], paddedBlock, os.FileMode(0644))
	if err != nil {
		log.Fatal(err)
	}

}

func pkcs7Pad(unpaddedBlock []byte, specificLength int) []byte {
	if len(unpaddedBlock) > specificLength {
		return []byte{}
	}

	paddedBlock := make([]byte, specificLength)

	var i int
	var b byte
	for i, b = range unpaddedBlock {
		paddedBlock[i] = b
	}

	for i++; i < specificLength; i++ {
		paddedBlock[i] = byte(0x04)
	}

	return paddedBlock
}
