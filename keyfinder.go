package main

/*
 * Part of cryptopals, Series 1, exercize 6,
 * find an XOR key based on given key length.
 */

import (
	"cryptopals/xor"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	keyBytes := flag.Int("k", 0, "key byte size guess")
	fileName := flag.String("f", "", "name of file containing potentially encrypted bytes")
	flag.Parse()

	if *keyBytes == 0 || *fileName == "" {
		fmt.Fprintf(os.Stderr, "Need key length (-k) and file name (-f) on command line\n")
		os.Exit(1)
	}

	ciphertext, err := ioutil.ReadFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	guessKeys(ciphertext, *keyBytes)
}

func guessKeys(ciphertext []byte, keylength int) {

	bins := binBytes(ciphertext, keylength)
	keybytes := make([]byte, keylength)
	candidates := make([]int, keylength)
	asciiCount := make([]int, keylength)

	for i := 0; i < keylength; i++ {
		keybytes[i], candidates[i], asciiCount[i] = findBestSingleByte(bins[i])
	}

	binsizes := make([]int, keylength)
	for i, bin := range bins {
		binsizes[i] = len(bin)
	}
	fmt.Printf("Bin sizes:      %v\n", binsizes)
	fmt.Printf("Best key bytes: %v\n", keybytes)
	fmt.Printf("Guesses:        %v\n", candidates) // indicates some keys got tried
	fmt.Printf("ASCII bytes:    %v\n", asciiCount)
	fmt.Printf("%q\n", string(keybytes))
}

func binBytes(ciphertext []byte, keylength int) [][]byte {

	bins := make([][]byte, keylength)

	for i, b := range ciphertext {
		binNo := i % keylength
		bins[binNo] = append(bins[binNo], b)
	}

	return bins
}

func findBestSingleByte(ciphertext []byte) (byte, int, int) {
	var bestKey byte
	var smallestTheta float64 = 1.0
	var asciiCount int
	var guessesMade int
	for keyVal := 0; keyVal < 256; keyVal++ {
		keyByte := byte(keyVal)
		if clearBytes, N := xor.CountAndXor(ciphertext, keyByte); N >= 9*len(ciphertext)/10 {
			guessesMade++
			clearVector := xor.NewAsciiByteVector(clearBytes)
			theta := xor.AsciiVectorAngle(clearVector, xor.AsciiLetterVector)
			if theta < smallestTheta {
				bestKey = keyByte
				smallestTheta = theta
				asciiCount = N
			}
		}
	}
	return bestKey, guessesMade, asciiCount
}
