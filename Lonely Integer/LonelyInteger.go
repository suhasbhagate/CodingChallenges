package main

import(
	"fmt"
)

func main(){
	arr := []int32{1, 2, 3, 4, 3, 2, 1}
	var res int32
	res = lonelyinteger(arr)
	fmt.Println(res)
}

func lonelyinteger(a []int32) int32 {
    m := make(map[int32]int32)
    for _, v := range a{
        m[v]++
    }
    var ans int32
    for k, v:= range m{
        if v ==1{
            ans = k
        }
    }
    return ans
}