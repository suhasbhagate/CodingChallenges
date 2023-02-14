package main

import(
	"fmt"
)

func main(){
	fmt.Println(kangaroo(0,3,4,2))
}


func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
    // Write your code here
    for{
        if x1+v1 == x2+v2{
            return "YES"
        }
        if x1>x2 && v1>v2{
            return "NO"
        }
        if x2>x1 && v2>v1{
            return "NO"
        }
        if x1==x2 && v1!=v2{
            return "NO"
        }
        if v1==v2 && x1!=x2{
            return "NO"
        }
        x1+=v1
        x2+=v2
    }
}