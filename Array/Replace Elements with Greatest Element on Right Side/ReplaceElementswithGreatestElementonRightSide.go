package main

import(
	"fmt"
)

func main(){
	arr := []int{17,18,5,4,6,1}
	fmt.Println(replaceElements(arr))
}


func replaceElements(arr []int) []int {
    max:= arr[len(arr)-1]
    for i:=len(arr)-2; i>=0; i--{
        ans:= max
        if max<arr[i]{
            max = arr[i]
        }
        arr[i] = ans
    }
    arr[len(arr)-1] = -1
    return arr
}