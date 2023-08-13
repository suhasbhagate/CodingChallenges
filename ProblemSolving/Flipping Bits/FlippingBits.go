package main

import (
	"fmt"
)

func main() {
	var n int64 = 9
	var flp int64
	flp = flippingBits(n)
	fmt.Println(flp)
}

func flippingBits(n int64) int64 {
	//var output uint32 = ^(uint32(n))
	//return int64(output)
	return int64(^uint32(n))
}
