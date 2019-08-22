package xor

import "math"

// Order of most frequent characters: " etaoinshrdlcumwfgypbvkjxqz"
// "non-alphabetic characters (digits, punctuation, etc.) collectively occupy the fourth position (having already included the space) between t and a."

var FrequencyIndex = map[rune]int{
	' ': 0,
	'e': 1,
	't': 2,
	// Punctuation (grouped) said to fit here
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

// English letter frequencies from Wikipedia
var EnglishLetterVector = map[byte]int{
	byte('a'): 8167,
	byte('b'): 1492,
	byte('c'): 2782,
	byte('d'): 4253,
	byte('e'): 12702,
	byte(' '): 13000,
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

// English ASCII character frequencies obtained
// by concatening many text files, then counting
// all the byte frequencies.
var AsciiLetterVector = map[byte]int{
	byte(0x09): 209,
	byte(0x0a): 2989,
	byte(0x20): 17105,
	byte(0x21): 17,
	byte(0x22): 276,
	byte(0x23): 53,
	byte(0x25): 3,
	byte(0x26): 13,
	byte(0x27): 59,
	byte(0x28): 353,
	byte(0x29): 370,
	byte(0x2a): 235,
	byte(0x2b): 26,
	byte(0x2c): 884,
	byte(0x2e): 1561,
	byte(0x2f): 793,
	byte(0x30): 310,
	byte(0x31): 228,
	byte(0x32): 187,
	byte(0x33): 92,
	byte(0x34): 61,
	byte(0x35): 51,
	byte(0x36): 66,
	byte(0x37): 59,
	byte(0x38): 82,
	byte(0x39): 73,
	byte(0x3a): 384,
	byte(0x3b): 10,
	byte(0x3c): 75,
	byte(0x3d): 402,
	byte(0x3e): 86,
	byte(0x3f): 7,
	byte(0x40): 8,
	byte(0x41): 348,
	byte(0x42): 177,
	byte(0x43): 344,
	byte(0x44): 246,
	byte(0x45): 323,
	byte(0x46): 97,
	byte(0x47): 296,
	byte(0x48): 86,
	byte(0x49): 480,
	byte(0x4a): 85,
	byte(0x4b): 12,
	byte(0x4c): 175,
	byte(0x4d): 143,
	byte(0x4e): 238,
	byte(0x4f): 207,
	byte(0x50): 292,
	byte(0x51): 4,
	byte(0x52): 286,
	byte(0x53): 393,
	byte(0x54): 425,
	byte(0x55): 149,
	byte(0x56): 54,
	byte(0x57): 77,
	byte(0x58): 88,
	byte(0x59): 55,
	byte(0x5a): 6,
	byte(0x5b): 307,
	byte(0x5c): 9,
	byte(0x5d): 309,
	byte(0x5e): 1,
	byte(0x5f): 304,
	byte(0x61): 5771,
	byte(0x62): 1861,
	byte(0x63): 3353,
	byte(0x64): 2887,
	byte(0x65): 9982,
	byte(0x66): 1769,
	byte(0x67): 2140,
	byte(0x68): 2663,
	byte(0x69): 6190,
	byte(0x6a): 290,
	byte(0x6b): 494,
	byte(0x6c): 4378,
	byte(0x6d): 2147,
	byte(0x6e): 5410,
	byte(0x6f): 6203,
	byte(0x70): 2120,
	byte(0x71): 78,
	byte(0x72): 5312,
	byte(0x73): 5498,
	byte(0x74): 7266,
	byte(0x75): 2649,
	byte(0x76): 999,
	byte(0x77): 928,
	byte(0x78): 292,
	byte(0x79): 1106,
	byte(0x7a): 90,
	byte(0x7b): 22,
	byte(0x7d): 22,
	byte(0x7e): 3,
	byte(0x84): 3,
	byte(0xa2): 3,
	byte(0xe2): 3,
}

// I think this one is incorrect: it does not use
// ' ' counts or any control characters ('\n', '\t')
// that might get counted.
func VectorAngle(vec1, vec2 map[byte]int) float64 {

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

func ByteVector1(cleartext []byte) map[byte]int {
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

func NewByteVector(cleartext []byte) map[byte]int {
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
func NewAsciiByteVector(cleartext []byte) map[byte]int {
	clearVector := make(map[byte]int)

	for _, b := range cleartext {
		if _, ok := clearVector[b]; !ok {
			clearVector[b] = 0
		}
		clearVector[b]++
	}

	return clearVector
}

func AsciiVectorAngle(vec1, vec2 map[byte]int) float64 {

	var sum1, sum2 int64
	var dotProduct int64
	for i := 0; i < 256; i++ {
		b := byte(i)
		c1, ok := vec1[b]
		if !ok {
			c1 = 0
		}
		c2, ok := vec2[b]
		if !ok {
			c2 = 0
		}
		dotProduct += int64(c1 * c2)
		sum1 += int64(c1 * c1)
		sum2 += int64(c2 * c2)
	}
	return math.Acos(float64(dotProduct) / math.Sqrt(float64(sum1*sum2)))
}
