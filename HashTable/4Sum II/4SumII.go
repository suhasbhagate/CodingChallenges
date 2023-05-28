package main

import(
	"fmt"
)


func main(){
	nums1 := []int{1,2}
	nums2 := []int{-2,-1}
	nums3 := []int{-1,2}
	nums4 := []int{0,2}
	fmt.Println(fourSumCount(nums1,nums2,nums3,nums4))
}


func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    var res int
    dp := make(map[int]int, len(nums1)*len(nums2))
    for _, a := range nums1 {
        for _, b := range nums2 {
            dp[a+b]++
        }
    }
    
    for _, c := range nums3 {
        for _, d := range nums4 {
            res += dp[0-c-d]
        }
    }
    return res
}