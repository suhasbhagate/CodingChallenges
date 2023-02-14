package main

import(
	"fmt"
)

func main(){
	var n int32 = 10
	var p int32 = 7
	var res int32
	res = pageCount(n, p)
	fmt.Println(res)
}

func pageCount(n int32, p int32) int32 {
    var end int32 = n/2
    var left int32 = p/2
    var right int32 = end - left
    var ans int32
    if left<right{
        ans = left
    } else{
        ans = right
    }
    return ans
}