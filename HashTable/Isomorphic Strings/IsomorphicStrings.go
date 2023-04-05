package main

import(
	"fmt"
	"strings"
)

func main(){
	s := "paper"
	t := "title"
	fmt.Println(isIsomorphic(s, t))
}

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    
    m1 := make(map[string]int)
    m2 := make(map[string]int)
    
    for _,v := range s{
        m1[string(v)]++
    }
    
    for _,v := range t{
        m2[string(v)]++
    }
    
    if len(m1) != len(m2){
        return false
    }    
    
    mp := make(map[string]string)
    for i:= range s{
        if _,ok := mp[string(s[i])]; !ok{
            mp[string(s[i])] = string(t[i])
        }        
    }
    
    var sb strings.Builder
    for j:= range s{
        sb.WriteString(mp[string(s[j])])
    }
    news:= sb.String()
    if news == t{
        return true
    } else{
        return false
    }
}