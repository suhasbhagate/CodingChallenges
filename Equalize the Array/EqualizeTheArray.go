package main

import(
	"fmt"
)

func main(){
	arr := []int32{1,2,2,3}
	fmt.Println(equalizeArray(arr))
}


func equalizeArray(arr []int32) int32 {
    m := make(map[int32]int32)
    for _,v:= range arr{
        m[v]++
    }
    
    var max int32 = 0
    var maxkey int32 = 0
    for k := range m{
        if m[k]>max{
            maxkey = k
            max = m[k]
        }
    }
    return int32(len(arr))-m[maxkey]
}