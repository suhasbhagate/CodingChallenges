package main

import(
	"fmt"
	"math"
)

func main(){
	nums :=[]int{2,2,3,1}
	fmt.Println(thirdMax(nums))
}


func thirdMax(nums []int) int {
    first:=math.MinInt
    second:=math.MinInt
    third:=math.MinInt
    for _,v:= range nums{
        if v>first{
            third = second
            second = first
            first = v
        } else if v>second && v <first{
            third = second
            second = v
        } else if v>third && v<second{
            third = v
        }
    }
    if third == math.MinInt{
        return first
    }
    return third
}