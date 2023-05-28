package main

import(
	"fmt"
)

func main(){
	s := "AAABBB"
	fmt.Println(alternatingCharacters(s))
}

func alternatingCharacters(s string) int32 {
    rns := []rune(s)
    var cnt int32 = 0
    for i:= 1; i<len(rns);i++{
        if rns[i]==rns[i-1]{
            cnt++
        }        
    }
    return cnt
}