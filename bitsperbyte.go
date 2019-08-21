package main

import "fmt"

func main() {
	for i := 0; i < 256; i++ {
		x := byte(i)
		bitcount := 0
		var j uint
		for j = 0; j < 8; j++ {
			bitcount += int((x >> j) & 0x01)
		}
		fmt.Printf("\t0x%02x: %d,\n", x, bitcount)
	}
}
