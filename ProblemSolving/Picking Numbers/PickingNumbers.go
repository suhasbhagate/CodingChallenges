package main

import(
	"fmt"
)

func main(){
	var a []int32 = []int32{4, 6, 5, 3, 3, 1}
	fmt.Println(pickingNumbers(a))
}

func pickingNumbers(a []int32) int32 {
    m:= make(map[int32]int32)
    for _,v:= range a{
        m[v]++        
    }
    var max int32 = 0
    for k,_:= range m{
        if m[k]+m[k+1]>max{
            max = m[k]+m[k+1]
        }
    }
    return max
}