package main

/*
Detect AES in ECB mode
In this file are a bunch of hex-encoded ciphertexts.

One of them has been encrypted with ECB.

Detect it.

Remember that the problem with ECB is that it is stateless and deterministic;
the same 16 byte plaintext block will always produce the same 16 byte
ciphertext.
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		if found, off1, off2 := detect(scanner.Text()); found {
			fmt.Printf("Line %d has duplicate substring, offsets %d and %d\n", lineCount, off1, off2)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}

func detect(s string) (bool, int, int) {

	b := []byte(s)

	limit := len(b) - 16

	for i := 0; i < limit; i += 16 {
		comparison := b[i : i+16]
		for j := i + 16; j < limit; j += 16 {
			if compareBuffers(comparison, b[j:j+16]) {
				return true, i, j
			}
		}
	}
	return false, -1, -1
}

func compareBuffers(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, ai := range a {
		if ai != b[i] {
			return false
		}
	}
	return true
}
