package main

import(
	"fmt"
)

func main(){
	nums := []int{3,3}
	fmt.Println(twoSum(nums,6))
}

func twoSum(nums []int, target int) []int {
    mp := make(map[int]int)
    for i,v := range nums{
        if val,ok := mp[target-v];ok{
            return []int{i,val}
        }
        mp[v] = i
    }
    return []int{-1,-1}
}