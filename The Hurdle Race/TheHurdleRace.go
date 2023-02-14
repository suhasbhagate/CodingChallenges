package main

import(
	"fmt"
)

func main(){
height := []int32{1, 6, 3, 5, 2}
var k int32 = 4
fmt.Println(hurdleRace(k,height))
}


func hurdleRace(k int32, height []int32) int32 {
    // Write your code here
    var max int32 = 0
    for _,v:= range height{
        if v>max{
            max = v
        }
    }
    if k>=max{
        return 0
    } else{
        return max-k
    }
}