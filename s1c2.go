package main

import (
	"bitsperbyte/xor"
	"encoding/hex"
	"fmt"
	"log"
)

/*

Fixed XOR

Write a function that takes two equal-length buffers and produces their XOR combination.

If your function works properly, then when you feed it the string:

1c0111001f010100061a024b53535009181c

... after hex decoding, and when XOR'd against:

686974207468652062756c6c277320657965

... should produce:

746865206b696420646f6e277420706c6179

*/

func main() {
	bufferStr := "1c0111001f010100061a024b53535009181c"
	keyStr := "686974207468652062756c6c277320657965"
	correctStr := "746865206b696420646f6e277420706c6179"

	buffer, err := hex.DecodeString(bufferStr)
	if err != nil {
		log.Fatal(err)
	}

	key, err := hex.DecodeString(keyStr)
	if err != nil {
		log.Fatal(err)
	}

	answer, err := FixedXOR(buffer, key)
	if err != nil {
		log.Fatal(err)
	}
	answerStr := hex.EncodeToString(answer)

	if answerStr == correctStr {
		fmt.Printf("passed\n")
		return
	}
	fmt.Printf("Failed\n")
}

// FixedXOR takes two equal-length buffers and produces their XOR combination.

func FixedXOR(b1, b2 []byte) ([]byte, error) {
	if len(b1) != len(b2) {
		return nil, fmt.Errorf("not same length buffers")
	}

	return xor.Encode(b1, b2), nil
}
