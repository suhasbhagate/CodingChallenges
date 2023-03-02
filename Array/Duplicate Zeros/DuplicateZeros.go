package main

import(
	"fmt"
)

func main(){
	arr := []int {1,0,2,3,0,4,5,0}
	duplicateZeros(arr)
}


func duplicateZeros(arr []int)  {
    i:=0
    for i<len(arr)-1{
		if arr[i]==0{
			for j:= len(arr)-2; j>i;j--{
				arr[j+1] = arr[j]
			}
			arr[i+1] = 0
			i++
		}
        i++
	}
	fmt.Println(arr)
}