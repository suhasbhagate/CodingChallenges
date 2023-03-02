package main

import(
	"fmt"
)

func main(){
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	//fmt.Println(removeDuplicates(nums))
    fmt.Println(removeDuplicatesArray(nums))
}


func removeDuplicates(nums []int) int {
    indx:=0
    for i:=1; i<len(nums); i++{
        if nums[indx] != nums[i]{
            indx++
            nums[indx] = nums[i]
        } 
    }
    return indx+1
}


func removeDuplicatesArray(nums []int)[]int{
    var ans []int
    ans = append(ans, nums[0])
    for i:=1;i<len(nums);i++{
        if nums[i]!=nums[i-1]{
            ans = append(ans, nums[i])
        }
    }
    return ans
}