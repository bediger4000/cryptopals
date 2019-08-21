package main

import (
	"bitsperbyte/xor"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	buffer, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	key := os.Args[2]
	zork := xor.Encode(buffer, []byte(key))
	outfilename := os.Args[3]
	err = ioutil.WriteFile(outfilename, zork, os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}
}
