package main

import(
	"fmt"
)

func main(){
	fmt.Println(andProduct(10, 14))
}


func andProduct(a int64, b int64) int64 {
    res := a
    for i:= a+1; i<=b; i+=2{
        res &= i
    }
    return res
}