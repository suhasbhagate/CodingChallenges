package main

import (
	"fmt"
	"math"
	"sort"
)

func main(){
	arr :=[]int32{1, 4, 7, 2}
	var k int32 = 2
	var res int32
	res = maxMin(k, arr)
	fmt.Println(res)
}

func maxMin(k int32, arr []int32) int32 {
    var slc []int
	for _, v := range arr{
		slc = append(slc, int(v))
	}
	sort.Ints(slc)
	fmt.Println(slc)
	var unfair float64 = math.MaxFloat64
	for i:= 0 ; i<=len(arr)-int(k); i++{
		low := slc[i]
		high := slc[i+int(k)-1]
		unfair = math.Min(unfair,float64(high)-float64(low))
	}
	return int32(unfair)
}