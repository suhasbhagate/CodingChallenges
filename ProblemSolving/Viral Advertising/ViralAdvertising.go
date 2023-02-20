package main

import(
	"fmt"
)

func main(){
	fmt.Println(viralAdvertising(3))
}

func viralAdvertising(n int32) int32 {
    shared := make([]int32,n)
    liked := make([]int32,n)
    cumulative := make([]int32,n)
    
    shared[0] = 5
    liked[0] = 2
    cumulative [0] = 2
    var i int32
    for i = 1; i<n; i++{
        shared[i] = liked[i-1] * 3
        liked[i] = shared[i]/2
        cumulative[i] = cumulative[i-1]+liked[i]
    }
    return cumulative[n-1]
}