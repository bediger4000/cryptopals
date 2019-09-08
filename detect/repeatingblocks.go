package detect

import "cryptopals/xor"

// DetectRepeatingBlocks returns bool (found a repeat if true)
// Should it just compare 1st blockLength bytes with 2nd in ciphertext array?
// Because we have control over input if this function is at all valuable.
func RepeatingBlocks(ciphertext []byte, blockLength int) bool {

	limit := len(ciphertext) - blockLength

	for offset1 := 0; offset1 < limit; offset1 += blockLength {
		originalBlock := ciphertext[offset1 : offset1+blockLength]
		for offset2 := offset1 + blockLength; offset2 < limit; offset2 += blockLength {
			comparisonBlock := ciphertext[offset2 : offset2+blockLength]
			if xor.CompareBuffers(originalBlock, comparisonBlock) {
				return true
			}
		}
	}
	return false
}
