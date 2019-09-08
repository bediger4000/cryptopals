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

	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	paddedBuffer := pkcs7.PadBlock(bytes, aes.BlockSize)

	if len(paddedBuffer)%aes.BlockSize != 0 {
		fmt.Fprintf(os.Stderr, "padded cleartext length %d is not a multiple of the block size\n", len(bytes))
	}

	dst := make([]byte, len(paddedBuffer))

	// all 0x00 initialization vector
	previousBlock := make([]byte, aes.BlockSize)

	for i := 0; i < len(bytes); i += aes.BlockSize {
		block.Encrypt(dst[i:i+aes.BlockSize], xor.Encode(paddedBuffer[i:i+aes.BlockSize], previousBlock))
		previousBlock = dst[i : i+aes.BlockSize]
	}

	base64text := base64.StdEncoding.EncodeToString(dst)

	_, err = os.Stdout.Write([]byte(base64text))
	if err != nil {
		log.Fatal(err)
	}
}
