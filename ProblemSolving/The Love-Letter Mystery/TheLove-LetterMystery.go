package main

import(
	"fmt"
)

func main(){
	s := "abcd"
	fmt.Println(theLoveLetterMystery(s))
}

func theLoveLetterMystery(s string) int32 {
    rns := []rune(s)
    var cnt int32 = 0
    i:= 0
    j:= len(s)-1
    for i<j{
        if rns[i]>='a' && rns[i]<='z' && rns[j]>='a' && rns[j]<='z' && rns[i] != rns[j]{
            if rns[i]<rns[j]{
                cnt += int32(rns[j]-rns[i])
            }
            if rns[i]>rns[j]{
                cnt += int32(rns[i]-rns[j])
            }
        }
        i++
        j--
    }
    return cnt
}