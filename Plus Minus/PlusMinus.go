package main

import(
	"fmt"
)

func main(){
	arr := []int32{-1, -1, 0, 1, 1}
	plusMinus(arr)
}

func plusMinus(arr []int32) {
    poscnt := 0
    negcnt := 0
    zerocnt := 0
    for _,v:= range arr{
        if v > 0{
            poscnt++
        } else if v < 0 {
            negcnt++
        } else {
            zerocnt++
        }
    }
    fmt.Printf("%.6f\n",float32(poscnt)/float32(len(arr)))
    fmt.Printf("%.6f\n",float32(negcnt)/float32(len(arr)))
    fmt.Printf("%.6f\n",float32(zerocnt)/float32(len(arr)))
}