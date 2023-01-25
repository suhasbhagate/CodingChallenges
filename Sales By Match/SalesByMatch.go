package main

import(
	"fmt"
)

func main(){
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	n := int32(len(ar))
	p := sockMerchant(n,ar)
	fmt.Println(p)
}

func sockMerchant(n int32, ar []int32) int32 {
    mp := make(map[int32]int)
    for _,v:= range ar{
        mp[v]++
    }
    cnt:=0
    for _, v:= range mp{
        for v>=2{
            v -=2
            cnt++
        }
    }
    return int32(cnt)
}