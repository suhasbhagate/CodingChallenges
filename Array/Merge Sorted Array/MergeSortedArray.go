package main

import(
	"fmt"
)

func main(){
	nums1:= []int{1,2,3,0,0,0}
	nums2:= []int{2,5,6}
	merge(nums1,3,nums2,3)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
    i:= 0
    j:= 0
    var res []int
    for i<m && j<n{
        if nums1[i]<=nums2[j]{
            res = append(res,nums1[i])
            i++
        } else{
            res = append(res,nums2[j])
            j++
        }
    }
    for i<m{
        res = append(res, nums1[i])
        i++
    }
    for j<n{
        res = append(res, nums2[j])
        j++
    }
    copy(nums1,res)
	fmt.Println(nums1)
}