package main

import(
	"fmt"
)

func main(){
	nums := []int{-1,0,3,5,9,12}
	//fmt.Println(search(nums,9))
	fmt.Println(RecursiveBinarySearch(nums,9,0,len(nums)-1))
}


func search(nums []int, target int) int {
    left:=0
    right:=len(nums)-1
    
    for(left<=right){
        mid := (left+right)/2
        if nums[mid]==target{
            return mid
        } else if target<nums[mid]{
            right = mid-1
        } else if target>nums[mid]{
            left = mid+1
        }
    }
    return -1
}

func RecursiveBinarySearch(nums []int, target, left, right int)int{
	mid := (left+right)/2
	if nums[mid]==target{
		return mid
	} else if target<nums[mid]{
		return RecursiveBinarySearch(nums,target,left,mid-1)
	} else if target>nums[mid]{
		return RecursiveBinarySearch(nums,target,mid+1,right)
	}
	return -1
}