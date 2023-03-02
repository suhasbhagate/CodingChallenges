package main

import(
	"fmt"
)

func main(){
	nums:= []int{-7,-3,2,3,11}
	fmt.Println(sortedSquares(nums))
}


func sortedSquares(nums []int) []int {
    var ans []int
    i:= 0
    for i< len(nums) && nums[i]<0{
        i++
    }
    m := i-1
    n := i
    
    for m>=0 && n<len(nums){
        one:= nums[m] * nums[m]
        two:= nums[n] * nums[n]
        if one<=two{
            ans = append(ans, one)
            m--
        } else{
            ans = append(ans, two)
            n++
        }
    }
    for m>=0{
        ans = append(ans, nums[m]*nums[m])
        m--
    }
    for n<len(nums){
        ans = append(ans, nums[n]*nums[n])
        n++
    }
    return ans
}