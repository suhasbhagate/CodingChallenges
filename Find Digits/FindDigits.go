package main

import(
	"fmt"
)

func main(){
	var n int32 = 124
	fmt.Println(findDigits(n))
}

func findDigits(n int32) int32 {
    var num int32 = n
    var slc []int32
    for num!= 0{
        slc = append(slc,num%10)
        num = num/10
    }
    var cnt int32 = 0
    for _,v:= range slc{
        if v!= 0 && n % v ==0{
            cnt++
        }
    }
    return cnt
}