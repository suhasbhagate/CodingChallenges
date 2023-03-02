package main

import(
	"fmt"
)

func main(){
	arr:= []int{0,3,2,1}
	fmt.Println(validMountainArray(arr))
}


func validMountainArray(arr []int) bool {
    if len(arr)<3{
        return false
    }
    i:=0
    for i<len(arr)-1 && arr[i]<arr[i+1]{
        i++
    }
    if i==0{
        return false
    }
    if i==len(arr)-1{
        return false
    }
    j:=i+1
    for j<len(arr) && arr[j]<arr[j-1]{
        j++
    }
    if j==len(arr){
        return true
    }
    return false
}