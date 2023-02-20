package main

import(
	"fmt"
)

func main(){
	a := []int32{-2, -1, 0, 1, 2}
	var k int32 = 3
	fmt.Println(angryProfessor())
}

func angryProfessor(k int32, a []int32) string {
    // Write your code here
    var cnt int32 = 0
    for _,v:= range a{
        if v<=0{
            cnt++
        }
    }
    if cnt>=k{
        return "NO"
    } else{
        return "YES"
    }    
}