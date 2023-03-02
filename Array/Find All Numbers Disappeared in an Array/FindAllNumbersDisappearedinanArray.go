package main

import(
	"fmt"
)

func main(){
	nums := []int{4,3,2,7,8,2,3,1}
	fmt.Println(findDisappearedNumbers(nums))
}

func findDisappearedNumbers(nums []int) []int {
    m := make(map[int]int)
    for _,v:= range nums{
        m[v]++
    }
    var res []int
    for i:=1; i<=len(nums);i++{
        if m[i]==0{
            res = append(res,i)
        }
    }
    return res
}