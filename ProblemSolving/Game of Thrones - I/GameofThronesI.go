package main

import(
	"fmt"
)


func main(){
	s := "aaabbbb"
	fmt.Println(gameOfThrones(s))
}


func gameOfThrones(s string) string {
    // Write your code here
    mp := make(map[rune]int)
    for _, v:= range s{
        mp[v]++
    }
    var ans string = "YES"
    if len(s) % 2 == 0{
        for k := range mp{
            if mp[k]%2 !=0{
                ans = "NO"
                break
            }
        }
    } else{
        cnt := 0
        for k:= range mp{
            if mp[k]%2 !=0{
                cnt++
            }
        }
        if cnt>1{
            ans = "NO"
        }
    }
    return ans
}