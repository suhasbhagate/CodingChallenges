package main

import(
	"fmt"
)

func main(){
	arr:= []int{9, -2, -9, 11, 56, -12, -3}
	fmt.Println(SquareEvenIndextedElement(arr))
}

func SquareEvenIndextedElement(arr []int)[]int{
	for i:=0; i<len(arr);i++{
		if i%2 ==0{
			arr[i] = arr[i]*arr[i]
		}
	}
	return arr
}