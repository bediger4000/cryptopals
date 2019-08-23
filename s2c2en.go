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

	// all 0x00 initialization vector
	previousBlock := make([]byte, aes.BlockSize)

	var i int
	for i = 0; i < len(bytes); i += aes.BlockSize {
		block.Encrypt(dst[i:i+aes.BlockSize], xor.Encode(bytes[i:i+aes.BlockSize], previousBlock))
		copy(previousBlock, dst[i:i+aes.BlockSize])
	}

	if i != len(bytes) {
		bytesEncrypted := i - aes.BlockSize
		fmt.Fprintf(os.Stderr, "Encrypted %d bytes\n", bytesEncrypted)
		paddedBytes := blocks.Pkcs7Pad(bytes[bytesEncrypted:], aes.BlockSize)
		block.Encrypt(dst[bytesEncrypted:], xor.Encode(paddedBytes[:], previousBlock))
	}

	base64text := base64.StdEncoding.EncodeToString(dst)

	_, err = os.Stdout.Write([]byte(base64text))
	if err != nil {
		log.Fatal(err)
	}
}
