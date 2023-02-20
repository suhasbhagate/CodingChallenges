package main

import(
	"fmt"
)

func main(){
	a:= []int32{3,4,5}
	var k int32 = 2
	queries := []int32{1,2}
	fmt.Println(circularArrayRotation(a,k,queries))
}

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
    // Write your code here
    var ans []int32
    k = k % int32(len(a))
    for _, v:= range queries{
        var indx int32 = v-k
        if indx<0{
            indx = indx + int32(len(a))
        }
        var res int32 = a[indx]
        ans = append(ans,res)
    }
    return ans
}