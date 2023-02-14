package main

import(
	"fmt"
)

func main(){
	var scores []int32 = []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	fmt.Println(breakingRecords(scores))
}


func breakingRecords(scores []int32) []int32 {
    // Write your code here
    var min int32 = scores[0]
    var max int32 = scores[0]
    var mincnt int32 = 0
    var maxcnt int32 = 0
    for _, v:= range scores{
        if v<min{
            min = v
            mincnt++
        } else if v>max{
            max = v
            maxcnt++
        }
    }
    return []int32{maxcnt,mincnt}
}