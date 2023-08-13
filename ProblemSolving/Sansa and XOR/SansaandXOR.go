package main

import(
	"fmt"
)

func main(){
	arr := []int32 {4,5,7,5}
	fmt.Println(sansaXor(arr))
}


func sansaXor(arr []int32) int32 {
    var res int32 = 0
    if len(arr) == 0 {
        res = 0
    } else if len(arr)==1{
        res = arr[0]
    } else if len(arr)%2 ==0{
        res = 0
    } else{
        for i:=0; i< len(arr); i+=2{
            res ^= arr[i]
        }
    }
    return res
}