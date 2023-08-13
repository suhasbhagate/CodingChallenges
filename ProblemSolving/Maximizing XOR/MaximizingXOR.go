package main

import (
	"fmt"
)


func main(){
	fmt.Println(maximizingXor(10,15))
}


func maximizingXor(l int32, r int32) int32 {
    var ans int32 = l ^ r
    var bitonepos int32 = 0
    for ans>0{
        bitonepos++
        ans>>=1
    }
    return (1<<bitonepos) - 1
}