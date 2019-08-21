package main

import (
	"bufio"
	"cryptopals/xor"
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

		var bestKey byte
		var smallestTheta float64 = 2.0

		for keyByte := 0; keyByte < 256; keyByte++ {
			clearBytes, n := xor.CountAndXor(ciphertext, byte(keyByte))
			if n == len(ciphertext) {
				clearVector := xor.NewByteVector(clearBytes)
				theta := xor.VectorAngle(clearVector, xor.EnglishLetterVector)
				if theta < smallestTheta {
					bestKey = byte(keyByte)
					smallestTheta = theta
				}
			}
		}

		if smallestTheta < 2.0 {
			fmt.Printf("line %d 0x%02x %f %q\n", lineCount, bestKey, smallestTheta, string(xor.Encode(ciphertext, []byte{bestKey})))
		}
	}

}
