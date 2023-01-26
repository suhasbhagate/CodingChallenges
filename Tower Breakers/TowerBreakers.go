package main

import(
	"fmt"
)

func main(){
	var n int32 = 4
	var m int32 = 6
	var res int32
	res = towerBreakers(n, m)
	fmt.Println(res)
}

func towerBreakers(n int32, m int32) int32 {
    if n%2 == 0 || m==1{
        return 2
    } else{
        return 1
    }
}