package main

import(
	"fmt"
)

func main(){
	nums := []int{3,1,2,4}
	//fmt.Println(sortArrayByParity(nums))
	fmt.Println(InPlaceSortArrayByParity(nums))
}


func sortArrayByParity(nums []int) []int {
    ans := make([]int,len(nums))
    m:=0
    n:= len(nums)-1
    for i:=0; i<len(nums);i++{
        if nums[i]%2 ==0{
            ans[m] = nums[i]
            m++
        } else{
            ans[n] = nums[i]
            n--
        }
    }
    return ans
}


func InPlaceSortArrayByParity(nums []int)[]int{
	m:= 0
    n:= len(nums)-1
    i:= 0
    for i<len(nums) && m<n{
        if nums[i]%2 ==0{
            if m!=i{
                nums[m],nums[i] = nums[i],nums[m]
                i--
            }
            m++
        } else{
            if n!= i{
                nums[n],nums[i] = nums[i], nums[n]
                i--
            }
            n--
        }
        i++
    }
    return nums
}