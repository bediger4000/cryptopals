package main

import (
	"cryptopals/myaes"
	"cryptopals/xor"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileName := flag.String("f", "", "name of file to encrypt, then decrypt")
	keyString := flag.String("k", "", "key string")
	flag.Parse()

	cleartext, err := ioutil.ReadFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes of clear text\n", len(cleartext))

	ciphertext := myaes.ECBEncode(cleartext, []byte(*keyString))
	fmt.Printf("%d bytes of cipher text\n", len(ciphertext))

	deciperedtext := myaes.ECBDecode(ciphertext, []byte(*keyString))
	fmt.Printf("%d bytes of deciphered text\n", len(deciperedtext))

	if xor.CompareBuffers(deciperedtext, cleartext) {
		fmt.Printf("Deciphered text matches cleartext\n")
		return
	}
	fmt.Printf("Deciphered text does not match cleartext\n")
}
