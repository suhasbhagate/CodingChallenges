package main

import(
	"fmt"
)

func main(){
	arr:= []int32{4, 2, 6, 1, 10}
	var n int32 = 5
	var k int32 = 3
	fmt.Println(workbook(n,k,arr))
}

func workbook(n int32, k int32, arr []int32) int32 {
    // Write your code here
    var page int32 = 0
    var cnt int32 = 0
    for _,v:= range arr{
        var indpg int32 = 0
        var st int32 = 0
        var end int32 = 0
        var data int32 = v
        for v>0{
            page++            
            st = indpg*k+1
            if v>=k{
                end = indpg*k+k
            } else{
                end = data
            }
            for i:=st; i<=end;i++{
                if i==page{
                    cnt++
                    //fmt.Printf("page: %d Chapter: %d Chapter Page: %d Problem: %d to %d\n",page,v,indpg, st,end)
                }
            }
            indpg++
            v-=k
        }
    }
    //fmt.Println(page)
    return cnt
}