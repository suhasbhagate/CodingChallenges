package main

import(
	"fmt"
)

func main(){
	var arr1 [3][3]int
	fmt.Println(arr1)

	arr2:= [3][3]int{}
	fmt.Println(arr2)

	var slc1 [][]int
	fmt.Println(slc1)

	slc2:= [][]int{}
	fmt.Println(slc2)

	slc3 := make([][]int,3)
	for i:= 0; i<3; i++{
		slc3[i] = make([]int,4)
	}
	fmt.Println(slc3)
}