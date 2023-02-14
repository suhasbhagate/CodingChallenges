package main

import(
	"fmt"
)

func main(){
	s:= "abcac"
	var n int64 = 10
	fmt.Println(repeatedString(s,n))
}


func repeatedString(s string, n int64) int64 {
    var original int64 = 0
    for _,v:= range s{
        if v=='a'{
            original++
        }
    }
    var full int64 = int64(n / int64(len(s)))
    var ans int64 = full * original
    var rem int64 = n % int64(len(s))
    var i int64
    for i =0; i<rem;i++{
        if s[i]=='a'{
            ans++
        }
    }
    return ans
}