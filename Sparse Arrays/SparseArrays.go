package main

import (
	"fmt"
)

func main(){
	strings := []string{"aba", "baba", "aba", "xzxb"}
	queries := []string{"aba", "xzxb", "ab"}
	res := matchingStrings(strings, queries)
	fmt.Println(res)
}

func matchingStrings(stringsArr []string, queries []string) []int32 {
    var arr []int32
    m := make(map[string]int32)
    for _,v := range stringsArr{
        m[v]++
    }

    for _, k := range queries{
        arr = append(arr, m[k])
    }
    return arr
}