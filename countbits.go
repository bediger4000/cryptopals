package main

import (
	"bitsperbyte/xor"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	bytes, err := hex.DecodeString(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	n := xor.CountBits(bytes)
	fmt.Printf("%d bits set in %q\n", n, os.Args[1])
}
