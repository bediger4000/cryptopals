package main

/*
 CBC mode is a block cipher mode that allows us to encrypt irregularly-sized
 messages, despite the fact that a block cipher natively only transforms
 individual blocks.

In CBC mode, each ciphertext block is added to the next plaintext block before
the next call to the cipher core.

The first plaintext block, which has no associated previous ciphertext block,
is added to a "fake 0th ciphertext block" called the initialization vector, or
IV.

Implement CBC mode by hand by taking the ECB function you wrote earlier, making
it encrypt instead of decrypt (verify this by decrypting whatever you encrypt
to test), and using your XOR function from the previous exercise to combine
them.

The file here is intelligible (somewhat) when CBC decrypted against "YELLOW
SUBMARINE" with an IV of all ASCII 0 (\x00\x00\x00 &c)

*/

import (
	"crypto/aes"
	"cryptopals/pkcs7"
	"cryptopals/xor"
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

	bytesLen := len(bytes)
	if bytesLen%aes.BlockSize != 0 {
		// It's ok to panic here, bytes[] should have length
		// a multiple of block size.
		panic("ciphertext is not a multiple of the block size")
	}

	dst := make([]byte, bytesLen)
	dummyBlock := make([]byte, aes.BlockSize) // Initialization vector
	lastBlock := make([]byte, aes.BlockSize)

	for i := 0; i < bytesLen; i += aes.BlockSize {
		block.Decrypt(dummyBlock, bytes[i:i+aes.BlockSize])
		copy(dst[i:i+aes.BlockSize], xor.Encode(dummyBlock, lastBlock))
		lastBlock = bytes[i : i+aes.BlockSize]
	}

	trimmed := pkcs7.TrimBuffer(dst)

	n, err := os.Stdout.Write(trimmed)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stderr, "Wrote %d bytes to stdout\n", n)
}
