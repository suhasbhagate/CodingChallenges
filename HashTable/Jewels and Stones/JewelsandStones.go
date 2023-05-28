package main

import(
	"fmt"
)

func main(){
	jewels := "aA"
	stones := "aAAbbbb"
	fmt.Println(numJewelsInStones(jewels,stones))
}


func numJewelsInStones(jewels string, stones string) int {
    mj := make(map[rune]int)
    
    rnsj := []rune(jewels)
    rnss := []rune(stones)
    
    for _,u:= range rnsj{
        mj[u]++
    }
    
    cnt :=0
    
    for _,v := range rnss{
        if _,ok := mj[v];ok{
            cnt++
        }
    }
    return cnt
}