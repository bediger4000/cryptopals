package xor

func Encode(buffer []byte, key []byte) []byte {
	keylen := len(key)
	ciphertext := make([]byte, len(buffer))
	for i, b := range buffer {
		ciphertext[i] = (b ^ key[i%keylen])
	}
	return ciphertext
}
