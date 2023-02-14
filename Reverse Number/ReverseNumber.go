package main

import(
	"fmt"
)

func main(){
	fmt.Println(Reverse(123))
}

func Reverse(n int32)int32{
    var rev int32 = 0
    for n!=0{
        rev = rev * 10
        rev = rev + n%10
        n = n/10
    }
    return rev
}