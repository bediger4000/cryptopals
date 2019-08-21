package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

/*
Convert hex to base64

The string:

49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d

Should produce:

SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t

So go ahead and make that happen. You'll need to use this code for the
rest of the exercises.

*/

func main() {
	str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	answer := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	plaintext, err := hex.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}

	b64 := encodeBase64Bytes(plaintext)
	fmt.Printf("%s\n", b64)

	if b64 == answer {
		fmt.Printf("Passed\n")
		return
	}
	fmt.Printf("Failed\n")
}

func encodeBase64Bytes(b []byte) string {

	return base64.StdEncoding.EncodeToString(b)
}
