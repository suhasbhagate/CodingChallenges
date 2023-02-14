package main

import(
	"fmt"
	"math"
)

func main(){
	//fmt.Println(Reverse(123))
	fmt.Println(beautifulDays(20,23,6))
}

func Reverse(n int32)int32{
    var rev int32 = 0
    for n!=0{
        rev = rev * 10
        rev = rev + n%10
        n = n/10
    }
    return rev
}

func beautifulDays(i int32, j int32, k int32) int32 {
    var cnt int32 = 0
    for a:=i; a<=j; a++{
        var diff int32= int32(math.Abs(float64(a) - float64(Reverse(a))))   
        if diff%k ==0{
            cnt++
        }        
    }
    return cnt
}