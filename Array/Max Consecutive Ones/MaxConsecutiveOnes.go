package main

import(
	"fmt"
)

func main(){
	nums := []int{1,1,0,1,1,1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}



func findMaxConsecutiveOnes(nums []int) int {  
    maxlen := 0
    cnt:= 0
    for i:= range nums{
        if nums[i]==1{
            cnt++
        } else{
            if cnt>maxlen{
                maxlen = cnt
            }
            cnt = 0
        }
    }
    if cnt>maxlen{
        maxlen = cnt
    }
    return maxlen
}