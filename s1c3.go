package main

import (
	"cryptopals/xor"
	"encoding/hex"
	"fmt"
	"log"
)

/*
Single-byte XOR cipher

The hex encoded string:

1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736

... has been XOR'd against a single character. Find the key, decrypt the message.

You can do this by hand. But don't: write code to do it for you.

How? Devise some method for "scoring" a piece of English plaintext. Character
frequency is a good metric. Evaluate each output and choose the one with the
best score.

*/

func main() {
	hexEncoded := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	ciphertext, err := hex.DecodeString(hexEncoded)
	if err != nil {
		log.Fatal(err)
	}

	var bestKey byte
	var smallestTheta float64 = 1.0

	for keyVal := 0; keyVal < 256; keyVal++ {
		keyByte := byte(keyVal)

		if clearBytes, N := xor.CountAndXor(ciphertext, keyByte); N == len(ciphertext) {
			clearVector := xor.NewByteVector(clearBytes)
			theta := xor.VectorAngle(clearVector, xor.EnglishLetterVector)
			if theta < smallestTheta {
				bestKey = keyByte
				smallestTheta = theta
				fmt.Printf("0x%02x  %f\n", bestKey, smallestTheta)
			}
		}
	}

	fmt.Printf("Best key byte 0x%02x: %q\n", bestKey, string(xor.Encode(ciphertext, []byte{bestKey})))

}
