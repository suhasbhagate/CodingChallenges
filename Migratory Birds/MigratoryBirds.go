package main

import(
	"fmt"
)

func main(){
	var arr []int32 = []int32{1, 4, 4, 4, 5, 3}
	fmt.Println(migratoryBirds(arr))
}

func migratoryBirds(arr []int32) int32 {
    // Write your code here
    m := make(map[int32]int32)
    for _,v:= range arr{
        m[v]++
    }
    var maxval int32 = 0
    var maxid int32
    for k,v:= range m{
        if v>maxval{
            maxval = v
            maxid = k
        }
        if v==maxval{
            if maxid>k{
                maxid = k
            }
            maxval = v
        }
    }
    return maxid
}