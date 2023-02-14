package main

import(
	"fmt"
	"sort"
)

func main(){
	var slc []int32 = []int32{5,30,24,22,13,44,33}
	fmt.Println(slc)
	sort.Slice(slc,func(i,j int)bool{return slc[i]<slc[j]})
	fmt.Println(slc)	
}