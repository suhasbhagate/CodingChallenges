package main

import(
	"fmt"
)

func main(){
	nums := []int{1,1,1,2,2,3}
	k := 2
	fmt.Println(topKFrequent(nums,k))
}

func topKFrequent(nums []int, k int) []int {
    mp := make(map[int]int)
    for _,v:= range nums{
        mp[v]++
    }
    var ans []int
    for i:=0; i<k; i++{
        max := 0
        maxind := 0
        for k:= range mp{
            if mp[k]>max{
                max= mp[k]
                maxind = k
            }
        }
        ans = append(ans,maxind)
        delete(mp,maxind)
    }
    return ans
}