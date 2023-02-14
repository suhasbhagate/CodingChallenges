package main

import(
	"fmt"
)

func main(){
	arr := []int32{2, 2, 1, 3, 2}
	var d int32 = 4
	var m int32 = 2
	var res int32
	res = birthday(arr, d, m)
	fmt.Println(res)
}

func birthday(s []int32, d int32, m int32) int32 {
    // Write your code here
    cnt := 0
    for i:= 0; i<=len(s)-int(m); i++{
        var sum int32 = 0
        for j:= i; j<i+int(m); j++{
            sum +=s[j]
        }
        if sum == d{
            cnt++
        }
    }
    return int32(cnt)
}