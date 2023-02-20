package main

import(
	"fmt"
)

func main(){
	var n int32 = 6
	var c int32 = 2
	var m int32 = 2
	fmt.Println(chocolateFeast(n,c,m))
}


func chocolateFeast(n int32, c int32, m int32) int32 {
    // Write your code here
    var cnt int32=0
    var wrapper int32= 0
    wrapper = n/c
    cnt = cnt+wrapper
    for wrapper>=m{
        var full int32 = wrapper/m
        cnt = cnt+full
        wrapper = wrapper - full*m
        wrapper = wrapper + full
    }
    return cnt
}