package pkcs7

/*
func main() {
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	blocksize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes, block size %d\n", N, blocksize)
	buffer := make([]byte, N)
	for i := 0; i < N; i++ {
		buffer[i] = byte(i)
	}
	zork := pkcs7padblock(buffer, blocksize)
	fmt.Printf("length of padded buffer %d\n", len(zork))

	fmt.Printf("%v\n", zork)

	doubleZork := pkcs7trimblock(zork)
	fmt.Printf("%v\n", doubleZork)
}
*/

// TrimBlock removes as much as 256 bytes off the tail of
// a buffer previously padded as per PKCS#7.
// The front of the input buffer comes back.
func TrimBuffer(buffer []byte) []byte {
	length := len(buffer)
	// PKCS#7 pad bytes value also count of pad bytes
	trimCount := int(buffer[length-1])
	return buffer[:length-trimCount]
}

// PadBlock composes a new buffer that's a multiple
// of blocksize, and in the case that input buffer is already
// a multiple of blocksize, is a whole block bigger.
// Final bytes (up to a whole block) comprise padding, with byte
// value of the number of padded bytes.
func PadBlock(buffer []byte, blocksize int) []byte {
	paddedBlockCount := len(buffer)/blocksize + 1
	spareByteCount := len(buffer) % blocksize
	fillInByteCount := blocksize - spareByteCount

	paddedBuffer := make([]byte, paddedBlockCount*blocksize)

	copy(paddedBuffer, buffer)

	// padBuffer a slice of paddedBuffer, so
	// paddedBuffer changes when padBuffer[i] get assigned to
	padBuffer := paddedBuffer[len(buffer):]
	for i := 0; i < fillInByteCount; i++ {
		padBuffer[i] = byte(fillInByteCount)
	}

	return paddedBuffer
}
