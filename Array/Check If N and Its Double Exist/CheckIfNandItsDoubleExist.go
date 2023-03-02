package main

import(
	"fmt"
)

func main(){
	arr:=[]int{10,2,5,3}
	fmt.Println(checkIfExist(arr))
}


func checkIfExist(arr []int) bool {
    if len(arr)<2{
        return false
    }
    for i:=1; i<len(arr);i++{
        for j:=0; j<i;j++{
            if arr[i]*2 == arr[j] || arr[i]==arr[j]*2{
                return true
            }
        }
    }
    return false
}