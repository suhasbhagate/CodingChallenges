package main

import(
	"fmt"
)

func main(){
	haystack:= "sadbutsad"
	needle := "sad"
	fmt.Println(strStr(haystack,needle))
}


func strStr(haystack string, needle string) int {
    if len(haystack)<len(needle){
		return -1
	}
	i:=0
    for i+len(needle)-1<len(haystack){
		if haystack[i:i+len(needle)]==needle{
			return i
		}
		i++
	}
	return -1
}