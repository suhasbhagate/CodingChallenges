package main

import(
	"fmt"
)

func main(){
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}


func firstUniqChar(s string) int {
    mp := make(map[string]int)
    for _,v := range s{
        mp[string(v)]++
    }
        
    for i,v := range s{
        if mp[string(v)]==1{
            return i
        }
    }
    return -1
}