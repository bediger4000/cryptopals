package main

import (
	"cryptopals/pkcs7"
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	padded := pkcs7.PadBlock([]byte(input), 16)
	fmt.Printf("Input:   %q\n", input)
	fmt.Printf("Padded:  %q\n", string(padded))
}
