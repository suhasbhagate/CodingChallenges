package main

import(
	"fmt"
)

func main(){
	nums := []int{1,2,3,4,5,6,7}
	fmt.Println(rotate(nums,3))
}


func rotate(nums []int, k int)[]int {
    arr := make([]int, len(nums))
    indx:=0
    for i:=0; i<len(nums);i++{
        indx = i-(k%len(nums))
        if indx<0{
            indx = indx+len(nums)
        }
        arr[i] = nums[indx]
    }
    copy(nums,arr)
	return nums
}