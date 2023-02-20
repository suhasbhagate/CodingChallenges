package main

import(
	"fmt"
)

func main(){
	c:= []int32{0, 0, 1, 0, 0, 1, 0}
	fmt.Println(jumpingOnClouds(c))
}


func jumpingOnClouds(c []int32) int32 {
    i := 0
    var jump int32
    for i <len(c)-1{
        if i<len(c)-2 && c[i+2]!= 1{
            i = i+2
            jump++
        } else{
            i++
            jump++
        }
        //fmt.Println(i)
    }
    return jump
}