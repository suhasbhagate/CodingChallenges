package main

import(
	"fmt"
	"sort"
)

func main(){
	slc := []int32{0, 1, 2, 4, 6, 5, 3}
	var med int32
	med = findMedian(slc)
	fmt.Println(med)
}

func findMedian(arr[] int32)int32{
	slc:=make([]int,len(arr))
	for _,v:= range arr{
		slc= append(slc,int(v))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(slc)))
	midx:=len(arr)/2
	med:=slc[midx]
	return int32(med)
}