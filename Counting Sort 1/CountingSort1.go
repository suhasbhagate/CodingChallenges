package main

import(
	"fmt"
)

func main(){
	arr := []int32{1, 1, 3, 2, 1}
	var res []int32
	res = countingSort(arr)
	fmt.Println(res)	
}

func countingSort(arr []int32) []int32 {
    result := make([]int32,100)
    for _, v := range arr{
        result[v]++
    }
    return result
}