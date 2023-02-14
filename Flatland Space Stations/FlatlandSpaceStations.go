package main

import(
	"fmt"
	"math"
)

func main(){
	var n int32 = 5
	c:= []int32{0, 4}
	fmt.Println(flatlandSpaceStations(n,c))
}

func flatlandSpaceStations(n int32, c []int32) int32 {
    dist:= make([]int32,n)
    var i int32 = 0
    for i=0;i<n;i++{
        var min int32 = math.MaxInt32
        for _,v:= range c{
            var md int32 = int32(math.Abs(float64(v)-float64(i)))
            if md<min{
                min = md
            }
        }
        dist[i]=min
    }
    var max int32= 0
    for _,u:= range dist{
        if u>max{
            max = u
        }
    }
    return max
}