package main

import(
	"fmt"
	"strings"
)

func main(){
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}


func reverseWords(s string) string {
    slc:= strings.Fields(s)
    var sb strings.Builder
    for i:=0;i<len(slc); i++{
        var sbl strings.Builder
        for j:=len(slc[i])-1; j>=0;j--{
            sbl.WriteString(string(slc[i][j]))
        }
        sb.WriteString(sbl.String())
        sb.WriteString(" ")
    }
    ans := sb.String()
    return ans[:len(ans)-1]
}