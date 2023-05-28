package main

import(
	"fmt"
)

func main(){
	s1 := "Hello"
	s2 := "World"
	fmt.Println(twoStrings(s1,s2))
}


func twoStrings(s1 string, s2 string) string {
    m1 := make(map[rune]int)
    m2 := make(map[rune]int)
    
    for _,i:= range s1{
        m1[i]++
    }
    for _,j := range s2{
        m2[j]++
    }
    var flg bool
    for k1:= range m1{
        if _,ok := m2[k1]; ok{
            flg = true
            break
        }
    }
    if flg == true{
        return "YES"
    } else {
        return "NO"
    }
}