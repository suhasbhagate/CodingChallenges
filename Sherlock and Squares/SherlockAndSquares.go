package main

import(
	"fmt"
	"math"
)

func main(){
	var a int32 = 9
	var b int32 = 17
	fmt.Println(squares(a,b)) 
}

func squares(a int32, b int32) int32 {
    num := int32(math.Sqrt(float64(a)))
    var cnt int32 = 0
    for num * num <=b{
        if num * num >=a{
            cnt++
        }
        num++
    }    
    return cnt
}