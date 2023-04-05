package main

import(
	"fmt"
)


func main(){
	fmt.Println(getRow(5))
}


func getRow(rowIndex int) []int {
    var ans []int
    prev:= 1
    ans = append(ans, prev)
    for i:=1; i<=rowIndex; i++{
        cur := (prev * (rowIndex - i +1))/i
        ans = append(ans, cur)
        prev = cur
    }
    return ans
}
