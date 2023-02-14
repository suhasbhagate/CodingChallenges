package main

import(
	"fmt"
	"math"
)

func main(){
	var n int64 = 5
	var res int64
	res = sumXor(n)
	fmt.Println(res)
}

func sumXor(n int64) int64 {
    var cnt int64= 0
    //var i int64
    // for i=0; i<=n; i++{
    //     //Solution 1
    //     // if (n + i) == (n ^ i){
    //     //     cnt++
    //     // }
        
    //     //Solution 2
    //     // if n & i == 0{
    //     //     cnt++
    //     // }
    // }
    
    //Count unset bits
    for n!= 0{
        if n & 1 ==0{
            cnt++
        }
        n >>= 1
    }
    res := math.Pow(2,float64(cnt))
    return int64(res)
}