package main

import(
	"fmt"
)

func main(){
	var n int32 = 5
	fmt.Println(utopianTree(n))
}

func utopianTree(n int32) int32 {
    slc := make([]int32,n+1)
    slc[0] = 1
    for i:=1; i<len(slc);i++{
        if i%2 == 0{
            slc[i] = slc[i-1]+1
        } else{
            slc[i] = slc[i-1]*2
        }        
    }
    return slc[n]
}