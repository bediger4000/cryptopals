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
	"bitsperbyte/xor"
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
	dummyBlock := make([]byte, aes.BlockSize)
	lastBlock := make([]byte, aes.BlockSize)

	for i := 0; i < len(bytes); i += aes.BlockSize {
		block.Decrypt(dummyBlock, bytes[i:i+aes.BlockSize])

		copy(dst[i:i+aes.BlockSize], xor.Encode(dummyBlock, lastBlock))
		copy(lastBlock, bytes[i:i+aes.BlockSize])
	}

	n, err := os.Stdout.Write(dst)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stderr, "Wrote %d bytes to stdout\n", n)
}
