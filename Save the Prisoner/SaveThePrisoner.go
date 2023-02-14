package main

import(
	"fmt"
)

func main(){
	var n int32 = 1
	var m int32 = 17894198
	var s int32 = 1
	fmt.Println(saveThePrisoner(n,m,s))
}

func saveThePrisoner(n int32, m int32, s int32) int32 {
    // Write your code here
    if n==1{
        return 1
    }
    m = m%n
    var ans int32 = (m+s)%n
    ans = ans - 1
    if ans < 1{
        ans = ans + n
    }
    return ans
}