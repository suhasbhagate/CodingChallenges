package main

import(
	"fmt"
	"math"
)

func main(){
	width := []int32{2, 3, 1, 2, 3, 2, 3, 3}
	cases := [][]int32{{0, 3}, {4, 6}, {6, 7}, {3, 5}, {0, 7}}
	var n int32 = 8
	fmt.Println(serviceLane(n,cases,width))
}


func serviceLane(n int32, cases [][]int32,width []int32) []int32 {
    // Write your code here
    var ans []int32
    for _,v:= range cases{
        st := v[0]
        end := v[1]
        var min int32 = math.MaxInt32
        for i:=st; i<=end;i++{
            if width[i]<min{
                min = width[i]
            }
        }
        ans = append(ans,min)
    }
    return ans
}