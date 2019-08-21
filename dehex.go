package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		line := scanner.Text()
		hexEncoded := strings.Trim(line, " \t\r\n")

		ciphertext, err := hex.DecodeString(hexEncoded)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(os.Stderr, "%d\n", len(ciphertext))
		fmt.Printf("%s", ciphertext)

	}
}
