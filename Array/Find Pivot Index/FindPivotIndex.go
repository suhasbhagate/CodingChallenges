package main

import(
	"fmt"
)

func main(){
	nums := []int{1,7,3,6,5,6}
	fmt.Println(pivotIndex(nums))
}


func pivotIndex(nums []int) int {
    sum:=0
    for _,v:= range nums{
        sum += v
    }
    leftsum:=0
    for i:= range nums{
        if leftsum == sum - leftsum - nums[i]{
            return i
        }
        leftsum+= nums[i]
    }
    return -1
}