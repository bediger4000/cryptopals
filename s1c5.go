package main

import (
	"encoding/hex"
	"fmt"
)

/*
 Here is the opening stanza of an important work of the English language:

Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal

Encrypt it, under the key "ICE", using repeating-key XOR.

In repeating-key XOR, you'll sequentially apply each byte of the key; the first byte of plaintext will be XOR'd against I, the next C, the next E, then I again for the 4th byte, and so on.

It should come out to:

0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f

*/

func main() {
	answer := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	textIn := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

	textOut := xorBytes([]byte(textIn), []byte("ICE"))

	candidate := hex.EncodeToString(textOut)
	fmt.Printf("%s\n", candidate)
	if candidate == answer {
		fmt.Printf("Passed\n")
	}
}

func xorBytes(text []byte, key []byte) []byte {
	returnText := make([]byte, len(text))
	lkey := len(key)

	for i, textByte := range text {
		returnText[i] = textByte ^ key[i%lkey]
	}
	return returnText
}
