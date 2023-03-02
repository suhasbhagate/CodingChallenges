package main

import(
	"fmt"
	"sort"
)

func main(){
	heights := []int{1,1,4,2,1,3}
	fmt.Println(heightChecker(heights))
}


func heightChecker(heights []int) int {
    expected:= make([]int,len(heights))
    copy(expected, heights)
    sort.Slice(expected,func (i, j int)bool{return expected[i]<expected[j]})
    cnt:= 0
    for i:= 0; i<len(heights);i++{
        if heights[i] != expected[i]{
            cnt++
        }
    }
    return cnt
}