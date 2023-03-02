package main

import(
	"fmt"
)

func main(){
	nums :=[]int{0,1,0,3,12}
	moveZeroes(nums)
}


func moveZeroes(nums []int)  {
    indx := 0
    for i:=0; i<len(nums);i++{
        if nums[i]!=0{
            nums[indx],nums[i] = nums[i],nums[indx]
            indx++
        }
    }
	fmt.Println(nums)
}