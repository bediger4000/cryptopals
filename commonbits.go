package main

import (
	"bitsperbyte/xor"
	"fmt"
	"os"
)

func main() {
	str := []byte(os.Args[1])
	key := []byte(os.Args[2])
	n := xor.CountBits(xor.Encode(str, key))
	fmt.Printf("%d bits in common\n", n)
}
