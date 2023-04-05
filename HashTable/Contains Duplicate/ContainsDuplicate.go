package main

import(
	"fmt"
)

func main(){
	nums := []int{1,1,1,3,3,4,3,2,4,2}
	fmt.Println(containsDuplicate(nums))
}


func containsDuplicate(nums []int) bool {
    m:= make(map[int]int)
    for _,v:= range nums{
        m[v]++
    }
    if len(m)==len(nums){
        return false
    } else{
        return true
    }
}