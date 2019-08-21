package main

import (
	"bitsperbyte/xor"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	buffer, err := hex.DecodeString(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	key, err := hex.DecodeString(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	zork := xor.Encode(buffer, key)
	fmt.Printf("%q\n", hex.EncodeToString(zork))
}
