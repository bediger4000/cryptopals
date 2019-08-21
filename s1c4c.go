package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	lineCount := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		lineCount++

		line := scanner.Text()
		hexEncoded := strings.Trim(line, " \t\r\n")

		ciphertext, err := hex.DecodeString(hexEncoded)
		if err != nil {
			log.Fatal(err)
		}

		for keyByte := 0; keyByte < 256; keyByte++ {
			clearBytes, n := xorAsciiBytes(ciphertext, byte(keyByte))
			if n == len(ciphertext) {
				fmt.Printf("%d %d 0x%02x %q\n", lineCount, n, keyByte, string(clearBytes))
			}
		}
	}

}

func xorAsciiBytes(ciphertext []byte, keyByte byte) ([]byte, int) {

	N := 0
	cleartext := make([]byte, len(ciphertext))
	for i, cipherByte := range ciphertext {
		clearByte := cipherByte ^ keyByte
		if clearByte >= 0x20 && clearByte <= 0x7e {
			N++
		} else {
			switch clearByte {
			case '\n':
				N++
			case '\t':
				N++

			}
		}
		cleartext[i] = clearByte
	}
	return cleartext, N
}
