package main

import (
	"cryptopals/detect"
	"cryptopals/myaes"
	"cryptopals/xor"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"strings"
)

var GlobalKey []byte

func main() {
	keyString := flag.String("k", "YELLOW SUBMARINE", "AES key")
	flag.Parse()

	// encodedString := `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK`
	// encodedString := `bm93IGlzIHRoZSB0aW1lIGZvciBhbGwgZ29vZCBtZW4K`
	// encodedString := `YWJjZGVmZ2hpamtsbW5vcAo=`
	encodedString := `YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXo=`

	unknownString, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("unknownString is %d bytes long\n", len(unknownString))

	blockLength := detect.FindCipherBlockLength(myaes.ECBEncode, []byte(*keyString))
	if blockLength > -1 {
		fmt.Printf("Cipher block length %d\n", blockLength)
	} else {
		fmt.Printf("No detectable cipher block length\n")
		return
	}

	findUnencryptedBytes(myaes.ECBEncode, blockLength, []byte(*keyString), unknownString)
}

func findUnencryptedBytes(cipher func([]byte, []byte) []byte, blockLength int, key []byte, unknownString []byte) {
	var decipheredBytes []byte

	fmt.Printf("Unknown string: %q\n", unknownString)
	encrypted := cipher(unknownString, key)
	blockCount := len(encrypted) / blockLength

	fmt.Printf("%d encrypted blocks\n", blockCount)

	//for blockNo := 0; blockNo < blockCount; blockCount++ {
	for i := blockLength; i > 0; i-- {
		var prefixBuffer []byte
		for j := 0; j < i-1; j++ {
			prefixBuffer = append(prefixBuffer, 'A')
		}
		// "AAAAAAA"
		fmt.Printf("Iter %d, prefix  %s\n", i, printBuffer(prefixBuffer))
		zork := make([]byte, len(prefixBuffer))
		copy(zork, prefixBuffer)
		zork = append(zork, unknownString...)
		fmt.Printf("Iter %d, combin  %s\n", i, printBuffer(zork))

		encryptedBytes := combineAndEncrypt(cipher, prefixBuffer, unknownString, key)
		encryptedBytes = encryptedBytes[0:blockLength]
		fmt.Printf("Encrypted buffr: %s\n", printBuffer(encryptedBytes))

		prefixBuffer = make([]byte, blockLength)
		for j := 0; j < i-1; j++ {
			prefixBuffer[j] = 'A'
		}
		for j := 0; j < len(decipheredBytes); j++ {
			prefixBuffer[i-1+j] = decipheredBytes[j]
		}
		fmt.Printf("Prefix buffer 2: %s\n", printBuffer(prefixBuffer))

		var candidate int
		for candidate = 0; candidate < 256; candidate++ {
			prefixBuffer[blockLength-1] = byte(candidate)
			fmt.Printf("Prefix buffer 3: %s\n", printBuffer(prefixBuffer))
			candidateBlock := cipher(prefixBuffer, key)
			fmt.Printf("candidate %3d:   %s\n", candidate, printBuffer(candidateBlock))
			if xor.CompareBuffers(candidateBlock[:blockLength], encryptedBytes) {
				fmt.Printf("Found candidate %3d, '%c'\n", candidate, candidate)
				decipheredBytes = append(decipheredBytes, prefixBuffer[blockLength-1])
				break
			}
		}
		if candidate == 256 {
			fmt.Printf("Did not find a byte that matched\n")
			break
		}
	}
	fmt.Printf("Decrypted: %q\n", string(decipheredBytes))
	//}
}

func combineAndEncrypt(cipher func([]byte, []byte) []byte, prefixString, unknownString, key []byte) []byte {
	return cipher(append(prefixString, unknownString...), key)
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
