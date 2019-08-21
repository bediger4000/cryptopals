package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math"
)

/*
Single-byte XOR cipher

The hex encoded string:

1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736

... has been XOR'd against a single character. Find the key, decrypt the message.

You can do this by hand. But don't: write code to do it for you.

How? Devise some method for "scoring" a piece of English plaintext. Character
frequency is a good metric. Evaluate each output and choose the one with the
best score.

*/

func main() {
	hexEncoded := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	ciphertext, err := hex.DecodeString(hexEncoded)
	if err != nil {
		log.Fatal(err)
	}

	var bestKey byte
	var smallestTheta float64 = 1.0

	for keyVal := 0; keyVal < 256; keyVal++ {
		keyByte := byte(keyVal)

		if xorAsciiBytes(ciphertext, keyByte) == len(ciphertext) {
			clearBytes := xorBytes(ciphertext, byte(keyByte))
			clearVector := byteVector(clearBytes)
			theta := vectorAngle(clearVector, englishReference)
			if theta < smallestTheta {
				bestKey = keyByte
				smallestTheta = theta
				fmt.Printf("0x%02x  %f\n", bestKey, smallestTheta)
			}
		}
	}

	fmt.Printf("Best key byte 0x%02x: %q\n", bestKey, string(xorBytes(ciphertext, bestKey)))

}

func xorAsciiBytes(ciphertext []byte, keyByte byte) int {

	N := 0
	for _, cipherByte := range ciphertext {
		clearByte := cipherByte ^ keyByte
		if clearByte >= 0x20 && clearByte <= 0x7e {
			N++
		}
	}
	return N
}

func xorBytes(ciphertext []byte, keyByte byte) []byte {

	clearByte := make([]byte, len(ciphertext))

	for i, cipherByte := range ciphertext {
		clearByte[i] = cipherByte ^ keyByte
	}

	return clearByte
}

// Order of most frequent characters: " etaoinshrdlcumwfgypbvkjxqz"
// "non-alphabetic characters (digits, punctuation, etc.) collectively occupy the fourth position (having already included the space) between t and a."

var frequencyIndex = map[rune]int{
	' ': 0,
	'e': 1,
	't': 2,

	'a': 4,
	'o': 5,
	'i': 6,
	'n': 7,
	's': 8,
	'h': 9,
	'r': 10,
	'd': 11,
	'l': 12,
	'c': 13,
	'u': 14,
	'm': 15,
	'w': 16,
	'f': 17,
	'g': 18,
	'y': 19,
	'p': 20,
	'b': 21,
	'v': 22,
	'k': 23,
	'j': 24,
	'x': 25,
	'q': 26,
	'z': 27,
}

var englishReference = map[byte]int{
	byte('a'): 8167,
	byte('b'): 1492,
	byte('c'): 2782,
	byte('d'): 4253,
	byte('e'): 12702,
	byte('f'): 2228,
	byte('g'): 2015,
	byte('h'): 6094,
	byte('i'): 6966,
	byte('j'): 153,
	byte('k'): 772,
	byte('l'): 4025,
	byte('m'): 2406,
	byte('n'): 6749,
	byte('o'): 7507,
	byte('p'): 1929,
	byte('q'): 95,
	byte('r'): 5987,
	byte('s'): 6327,
	byte('t'): 9056,
	byte('u'): 2758,
	byte('v'): 978,
	byte('w'): 2360,
	byte('x'): 150,
	byte('y'): 1974,
	byte('z'): 74,
}

func vectorAngle(vec1, vec2 map[byte]int) float64 {

	var sum1, sum2 int64
	var dotProduct int64
	for keyByte := 'a'; keyByte <= 'z'; keyByte++ {
		b := byte(keyByte)
		dotProduct += int64(vec1[b] * vec2[b])
		sum1 += int64(vec1[b] * vec1[b])
		sum2 += int64(vec2[b] * vec2[b])
	}
	return math.Acos(float64(dotProduct) / math.Sqrt(float64(sum1*sum2)))
}

func byteVector1(cleartext []byte) map[byte]int {
	clearVector := make(map[byte]int)

	for _, b := range cleartext {
		if _, ok := clearVector[b]; !ok {
			if b >= 0x41 && b <= 0x5a {
				b += 0x20
			}
			clearVector[b] = 0
		}
		clearVector[b]++
	}

	return clearVector
}

func byteVector(cleartext []byte) map[byte]int {
	clearVector := make(map[byte]int)

	for _, b := range cleartext {
		if (b == ' ') || (b >= 0x41 && b <= 0x5a) || (b >= 0x61 && b <= 0x7a) {
			if _, ok := clearVector[b]; !ok {
				if b >= 0x41 && b <= 0x5a {
					b += 0x20
				}
				clearVector[b] = 0
			}
			clearVector[b]++
		}
	}

	return clearVector
}
