package detect

// Find block length of a cipher by increasing input buff length
// by 1 and seeing where the output steps up in lenghth.
func FindCipherBlockLength(cipher func([]byte, []byte) []byte, key []byte) int {

	var cipherTextLengths []int
	var lastCipherTextLength int

	for bufferLength := 1; bufferLength < 128; bufferLength++ {
		var buffer []byte

		for i := 0; i < bufferLength; i++ {
			buffer = append(buffer, 'A')
		}
		ciphertext := cipher(buffer, key)
		if len(ciphertext) != lastCipherTextLength {
			cipherTextLengths = append(cipherTextLengths, len(ciphertext))
		}
		lastCipherTextLength = len(ciphertext)
		// fmt.Printf("buffer length %d, ciphertext length %d\n", bufferLength, len(ciphertext))
		// fmt.Printf("Ciphertext: %v\n", ciphertext)
	}

	// fmt.Printf("Cipher text lengths: %v\n", cipherTextLengths)

	blocksize := cipherTextLengths[0]

	if len(cipherTextLengths) > 1 {
		blocksize = gcd(cipherTextLengths[0], cipherTextLengths[1])
	}

	return blocksize
}

// Greatest Common Denominator, Euclidean Algorithm
func gcd(a, b int) int {
	for a != b {
		for a > b {
			a -= b
		}
		for b > a {
			b -= a
		}
	}

	return a
}
