package main

import(
	"fmt"
	"strings"
	"regexp"
)

func main(){
	str := "We promptly judged antique ivory buckles for the next prize"
	fmt.Println(pangrams(str))
}


func pangrams(s string) string {
    var result string
    str := strings.ToUpper(s)
    str = regexp.MustCompile(`[^A-Z]`).ReplaceAllString(str,"")
    rns := []rune(str)
        
    if len(rns)<26{
        result = "not pangram"
    } else{
        m := make(map[rune]bool)
        for _, v:= range rns{
            m[v] = true
        }
        if len(m)==26{
            result = "pangram"
        } else{
            result = "not pangram"
        }
    }    
return result
}