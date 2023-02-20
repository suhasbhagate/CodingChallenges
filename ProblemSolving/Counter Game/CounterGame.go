package main

import (
	"fmt"
	"math"
)

func main(){
	var n int64= 132
	res := counterGame(n)
	fmt.Println(res)
}

func counterGame(n int64) string {
	cnt := 0
    for n>1{
		p := int64(math.Log2(float64(n)))
		
        if n == int64(math.Pow(2,float64(p))){
            n = n/2
        } else{
            n = n - int64(math.Pow(2,float64(p)))
        }
		cnt++
    }
    if cnt % 2 ==1{
        return "Louise"
    } else{
        return "Richard"
    }
}
