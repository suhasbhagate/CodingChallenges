package main

import(
	"fmt"
)

func main(){
	var n int32 = 2
	queries:= [][]int32{{1, 0, 5}, {1, 1, 7}, {1, 0, 3}, {2, 1, 0}, {2, 1, 1}}
	var res []int32
	res = dynamicArray(n, queries)
	fmt.Println(res)
}

func dynamicArray(n int32, queries [][]int32) []int32 {
    var result []int32
    arr := make([][]int32, n)
    var lastAnswer int32 = 0
    for _,v:= range queries{
        querytype:= v[0]
        x := v[1]
        y := v[2]
        var idx int32
        idx = ((x^lastAnswer)%n)
        if querytype==1{
           arr[idx] = append(arr[idx],y)
        } else if querytype==2{
           lastAnswer = arr[idx][y % int32(len(arr[idx]))]
           result = append(result,lastAnswer)
        }
    }
    return result
}