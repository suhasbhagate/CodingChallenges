package main

import(
	"fmt"
	"sort"
)


func main(){
	nums :=[]int{6,2,6,5,1,2}
	fmt.Println(arrayPairSum(nums))
}


func arrayPairSum(nums []int) int {
    sort.Slice(nums,func(i,j int)bool{return nums[i]<nums[j]})
    ans:=0
    for i:=0; i<len(nums);i=i+2{
        ans+=nums[i]
    }
    return ans
}