package main

import(
	"fmt"
	"strconv"
)

func main(){
	nums:= []int{12,345,2,6,7896}
	fmt.Println(findNumbers(nums))
}


func findNumbers(nums []int) int {
    cnt:=0
    for _,v:= range nums{
        if len(strconv.Itoa(v))%2==0{
            cnt++
        }
    }
    return cnt
}