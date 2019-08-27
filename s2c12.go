package main

import (
	"cryptopals/detect"
	"cryptopals/myaes"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"strings"
)

var GlobalKey []byte

func main() {
	keyString := flag.String("k", "YELLOW SUBMARINE", "AES key")
	myString := flag.String("s", "", "prefix string")
	flag.Parse()

	prefixString := []byte(*myString)

	// encodedString := `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK`
	encodedString := `bm93IGlzIHRoZSB0aW1lIGZvciBhbGwgZ29vZCBtZW4K`

	unknownString, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("unknownString is %d bytes long\n", len(unknownString))
	fmt.Printf("prefix        is %d bytes long\n", len(prefixString))

	blockLength := detect.FindCipherBlockLength(myaes.ECBEncode, []byte(*keyString))
	if blockLength > -1 {
		fmt.Printf("Cipher block length %d\n", blockLength)
	} else {
		fmt.Printf("No detectable cipher block length\n")
		return
	}

	findFirstUnencrypteByte(myaes.ECBEncode, blockLength, []byte(*keyString), unknownString)
}

func findFirstUnencrypteByte(cipher func([]byte, []byte) []byte, blockLength int, key []byte, unknownString []byte) {
	// prefixLength := blockLength - 1
	prefixBuffer := make([]byte, blockLength-1)
	for i := 0; i < blockLength-1; i++ {
		prefixBuffer[i] = 'A'
	}
	fmt.Printf("prefix block:                 %s\n", printBuffer(prefixBuffer))

	cipherText := combineAndEncrypt(prefixBuffer, unknownString, key)
	searchValue := cipherText[blockLength-1]
	fmt.Printf("Searching for byte with encrypted value 0x%02x\n", cipherText[blockLength-1])
	fmt.Printf("Block to look for:            %s\n", printBuffer(cipherText[:blockLength]))

	prefixBuffer = make([]byte, blockLength)
	for i := 0; i < blockLength-1; i++ {
		prefixBuffer[i] = 'A'
	}

	for i := 0; i < 256; i++ {
		prefixBuffer[blockLength-1] = byte(i)
		cipherText := cipher(prefixBuffer, key)
		fmt.Printf("input byte 0x%02x, cipher block %s\n", byte(i), printBuffer(cipherText[:blockLength]))
		if cipherText[blockLength-1] == searchValue {
			fmt.Printf("Found searchvalue 0x%x, '%c' plaintext\n", searchValue, prefixBuffer[blockLength-1])
		}
	}
}

func combineAndEncrypt(prefixString, unknownString, key []byte) []byte {
	return myaes.ECBEncode(append(prefixString, unknownString...), key)
}

func printBuffer(buffer []byte) string {
	sb := &strings.Builder{}
	divider := ""
	for _, b := range buffer {
		sb.WriteString(fmt.Sprintf("%s%3d", divider, b))
		divider = " "
	}
	return sb.String()
}
