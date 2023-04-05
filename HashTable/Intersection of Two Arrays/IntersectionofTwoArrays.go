package main

import(
	"fmt"
)

func main(){
	nums1 := []int{4,9,5}
	nums2 := []int{9,4,9,8,4}
	fmt.Println(intersection(nums1,nums2))
}


func intersection(nums1 []int, nums2 []int) []int {
    m1 := make(map[int]int)
    m2 := make(map[int]int)
    
    for _,v := range nums1{
        m1[v]++
    }
    
    for _,v := range nums2{
        m2[v]++
    }
    
    var ans []int
    for k := range m1{
        if _,ok := m2[k];ok{
            ans = append(ans, k)
        }
    }
    return ans
}