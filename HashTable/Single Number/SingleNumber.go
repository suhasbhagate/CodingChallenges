package main

import(
	"fmt"
)

func main(){
	nums :=[]int{4,1,2,1,2}
	fmt.Println(singleNumber(nums))
}


func singleNumber(nums []int) int {
    m := make(map[int]int)
    for _,v := range nums{
        m[v]++
    }
    for k:= range m{
        if m[k]==1{
            return k
        }
    }
    return -1
}