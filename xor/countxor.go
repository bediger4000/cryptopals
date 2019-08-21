package xor

func CountAndXor(ciphertext []byte, keyByte byte) ([]byte, int) {
	N := 0
	cleartext := make([]byte, len(ciphertext))
	for i, cipherByte := range ciphertext {
		clearByte := cipherByte ^ keyByte
		if clearByte >= 0x20 && clearByte <= 0x7e {
			N++
		} else if clearByte == '\n' || clearByte == '\t' {
			N++
		}
		cleartext[i] = clearByte
	}
	return cleartext, N
}
