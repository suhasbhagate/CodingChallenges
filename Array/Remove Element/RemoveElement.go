package main

import(
	"fmt"
)

func main(){
 nums := []int {3,2,2,3}
 fmt.Println(removeElement(nums,3))
}


func removeElement(nums []int, val int) int {
    indx := 0
    for i:=0;i<len(nums);i++{
        if nums[i]!= val{
            nums[indx] = nums[i]
            indx++
        }
    }
    return indx
}