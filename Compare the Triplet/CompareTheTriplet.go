package main

import(
	"fmt"
)

func main(){
	var a []int32 = []int32{5, 6, 7}
	var b []int32 = []int32{3, 6, 10}
	fmt.Println(compareTriplets(a,b))
}

func compareTriplets(a []int32, b []int32) []int32 {
    var acnt int32 = 0
    var bcnt int32 = 0
    for i:=0; i<len(a);i++{
        if a[i]>b[i]{
            acnt++
        } else if a[i]<b[i]{
                bcnt++
        }
    }
    return []int32{acnt,bcnt}
}