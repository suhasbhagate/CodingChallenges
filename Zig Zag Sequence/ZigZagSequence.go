package main

import(
	"fmt"
	"sort"
)

func main(){
	arr:= []int{2, 3, 5, 1, 4}
	var res []int
	res = findZigZagSequence(arr)
	fmt.Println(res)
}


func findZigZagSequence(arr []int)[]int{
	sort.Ints(arr)
	mid := (len(arr)-1)/2
	for i,j:=mid,len(arr)-1; i<j; i,j = i+1,j-1{
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}