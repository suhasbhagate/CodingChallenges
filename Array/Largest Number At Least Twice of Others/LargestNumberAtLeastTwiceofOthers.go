package main

import(
	"fmt"
)

func main(){
	nums:= []int{3,6,1,0}
	fmt.Println(dominantIndex(nums))
}


func dominantIndex(nums []int) int {
    max:= nums[0]
    indx:=0
    for i:=1; i<len(nums);i++{
        if nums[i]>max{
            max = nums[i]
            indx = i
        }
    }
    for i:=0; i<len(nums);i++{
        if i!=indx{
            if max<2*nums[i]{
                return -1
            }
        }
    }
    return indx
}