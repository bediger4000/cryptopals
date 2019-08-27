package main

import (
	"cryptopals/myaes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	debug := flag.Bool("x", false, "debug output on stderr")
	decode := flag.Bool("d", false, "decode input")
	encode := flag.Bool("e", false, "encode input")
	fileName := flag.String("f", "", "name of file to encrypt, then decrypt")
	keyString := flag.String("k", "", "key string")
	flag.Parse()

	if !*decode && !*encode {
		log.Fatal("Need one of -d or -e on command line\n")
	}

	if *fileName == "" {
		*fileName = os.Args[len(os.Args)-1]
	}

	if *debug {
		fmt.Fprintf(os.Stderr, "Reading in file named %q\n", *fileName)
	}

	text, err := ioutil.ReadFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	if *debug {
		fmt.Fprintf(os.Stderr, "Read in %d bytes\n", len(text))
	}

	var outputtext []byte

	if *encode {
		outputtext = myaes.ECBEncode(text, []byte(*keyString))

		if *debug {
			fmt.Fprintf(os.Stderr, "%d bytes of ciphertext\n", len(outputtext))
		}
	}

	if *decode {
		outputtext = myaes.ECBDecode(text, []byte(*keyString))

		if *debug {
			fmt.Fprintf(os.Stderr, "%d bytes of cleartext\n", len(outputtext))
		}
	}

	os.Stdout.Write(outputtext)
}
