package main

import(
	"fmt"
)

func main(){
	fmt.Println(howManyGames(20,3,6,70))
}

func howManyGames(p int32, d int32, m int32, s int32) int32 {
    // Return the number of games you can buy
    var bill int32 = 0
    var cnt int32 = 0

    if p>s{
        return 0
    }
    
    if p<= s{
        cnt++
        bill = bill+p
    }
    var cost int32 = p-d
    for bill<=s && cost>m{
        bill = bill+cost
        cnt++
        cost = cost - d
    }
    cnt--
    for bill<=s{
        cnt++
        bill = bill+m
    }
    return cnt
}