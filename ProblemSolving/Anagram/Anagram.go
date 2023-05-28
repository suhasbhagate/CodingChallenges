package main

import(
	"fmt"
)

func main(){
	s := "aaabbb"
	fmt.Println(anagram(s))
}


func anagram(s string) int32 {
    // Write your code here
    if len(s)%2 !=0{
        return -1
    }
    rns := []rune(s)
    mp := make(map[rune]int32)

    for i:= 0; i<len(s)/2; i++{
        mp[rns[i]]++
    }
    var res int32 = 0
    for j:=len(s)/2; j<len(s);j++{
        if mp[rns[j]]>0{
            mp[rns[j]]--
        }
    }   
    
    for k:= range mp{
        res += mp[k]
    }
    return res
}