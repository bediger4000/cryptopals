package main

import (
	"bitsperbyte/xor"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	minKeyBytes := flag.Int("m", 2, "minimum key guess byte size")
	maxKeyBytes := flag.Int("M", 20, "maximum key guess byte size")
	fileName := flag.String("f", "", "name of file containing potentially encrypted bytes")
	flag.Parse()

	bytes, err := ioutil.ReadFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	tryKeySizes(bytes, *minKeyBytes, *maxKeyBytes)
}

func tryKeySizes(buffer []byte, minKeyLength, maxKeyLength int) {
	for keyLength := minKeyLength; keyLength <= maxKeyLength; keyLength++ {
		sumBits, bufferCount := compareSubBuffers(keyLength, buffer)
		hammingDistance := float64(sumBits) / float64(bufferCount*keyLength)
		fmt.Printf("%d\t%.04f\n", keyLength, hammingDistance)
	}
}

func compareSubBuffers(keyLength int, buffer []byte) (int, int) {
	bufferLength := len(buffer)
	bufferComparisons := 0
	commonBitsCount := 0
	comparisonBuffer := buffer[0:keyLength]
	for j := keyLength + 1; j < bufferLength; j += keyLength {
		xoredSubBuffer := xor.Encode(comparisonBuffer, buffer[j:j+keyLength])
		commonBitsCount += xor.CountBits(xoredSubBuffer)
		bufferComparisons++
	}
	return commonBitsCount, bufferComparisons
}
