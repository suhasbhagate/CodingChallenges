package main

import(
	"fmt"
	"sort"
)

func main(){
	arr := []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr)
}


func miniMaxSum(arr []int32) {
    var slc []int
    for _,v := range arr{
        slc = append(slc,int(v))        
    }    
    sort.Ints(slc)
    var minsum int64 = 0
    var maxsum int64 = 0
    for i, v := range slc{
        if i != 0{
            maxsum += int64(v)
        }
        if i != len(arr)-1{
            minsum += int64(v)
        }
    }
    fmt.Printf("%d %d", minsum, maxsum)
}