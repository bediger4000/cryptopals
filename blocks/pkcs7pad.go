package blocks

//
func Pkcs7Pad(unpaddedBlock []byte, specificLength int) []byte {
	if len(unpaddedBlock) > specificLength {
		return []byte{}
	}

	paddedBlock := make([]byte, specificLength)

	var i int
	var b byte
	for i, b = range unpaddedBlock {
		paddedBlock[i] = b
	}

	for i++; i < specificLength; i++ {
		paddedBlock[i] = byte(0x04)
	}

	return paddedBlock
}
