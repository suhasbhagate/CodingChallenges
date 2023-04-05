package main

import(
	"fmt"
)

func main(){
	nums := []int{1,2,3,1,2,3}
	k := 2
	fmt.Println(containsNearbyDuplicate(nums,k))
}


func containsNearbyDuplicate(nums []int, k int) bool {
    mp := make(map[int]int)
    for _,v := range nums{
        mp[v]++
    }
    for key := range mp{
        if mp[key]>1{
            for st:=0; st<len(nums)-1;st++{
                if nums[st]==key{
                    end := st+1
                    for end<len(nums) && end<=st+k{
                        if nums[end]==key{
                            return true
                        }
                        end++
                    }
                }
            }
            
        }
    }
    return false
}