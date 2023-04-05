package main

import(
	"fmt"
)

func main(){
	nums1 := []int{4,9,5}
	nums2 := []int{9,4,9,8,4}
	fmt.Println(intersect(nums1,nums2))
}


func intersect(nums1 []int, nums2 []int) []int {
    m1 := make(map[int]int)
    m2 := make(map[int]int)
    for _,v := range nums1{
        m1[v]++
    }
    for _,v := range nums2{
        m2[v]++
    }
    
    var ans []int
    for k, v1 := range m1{
        cnt := 0
        if v2,ok := m2[k]; ok{
            if v1<v2{
                cnt = v1
            } else{
                cnt = v2
            }
            for j:=0; j<cnt; j++{
                ans = append(ans, k)
            }
        }
    }
    return ans
}