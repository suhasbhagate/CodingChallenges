package main

import(
	"fmt"
	"math"
)


func main(){
	nums := []int{2,3,1,2,4,3}
	target := 7
	fmt.Println(minSubArrayLen(target,nums))
}


func minSubArrayLen(target int, nums []int) int {
    st:=0
    end:=0
    sum:=0
    minlen:=math.MaxInt
    flg:=false
    for end<len(nums){
        sum+=nums[end]
        if sum>=target{
            flg = true
            if end-st+1< minlen{
                minlen = end-st+1
            }
            for sum>=target{
                sum = sum-nums[st]                
                if end-st+1< minlen{
                    minlen = end-st+1
                }
                st++
            }            
        }
        end++
    }
    if flg==true{
        return minlen
    } else{
        return 0
    }
}