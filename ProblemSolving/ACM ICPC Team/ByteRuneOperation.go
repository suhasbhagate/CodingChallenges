package main

import(
	"fmt"
)

func main(){
	str := "abcd"
	rns := []rune(str)
	fmt.Println(rns)
	fmt.Println(rns[01]-rns[0])

	bts := []byte(str)
	fmt.Println(bts)
	fmt.Println(bts[01]-bts[0])
}