package main

import(
	"fmt"
)

func main(){
	str := "abca"
	fmt.Println(stringConstruction(str))
}


func stringConstruction(s string) int32 {
    mp := make(map[rune]bool)
    for _,v := range s{
        mp[v] = true
    }
    return int32(len(mp))
}