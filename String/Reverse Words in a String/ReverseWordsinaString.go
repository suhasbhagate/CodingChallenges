package main

import(
	"fmt"
	"strings"
)

func main(){
	s:="the sky is blue"
	fmt.Println(reverseWords(s))
}


func reverseWords(s string) string {
    slc := strings.Fields(s)
    var sb strings.Builder
    for i:=len(slc)-1; i>0; i--{
        sb.WriteString(slc[i])
        sb.WriteString(" ")
    }
    sb.WriteString(slc[0])
    return sb.String()
}