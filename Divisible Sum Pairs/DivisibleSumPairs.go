package main

import(
	"fmt"
)

func main(){
	var ar []int32 = []int32{1, 3, 2, 6, 1, 2}
	var k int32 = 3
	fmt.Println(divisibleSumPairs(6,k,ar))
}


func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
    // Write your code here
    var cnt int32 = 0
    var n1, n2 int32
    for i:=0; i<len(ar)-1;i++{
        n1 = ar[i]
        for j:=i+1; j<len(ar);j++{
            n2= ar[j]
            if (n1+n2) % k ==0{
                cnt++
                
            }
        }
    }
    return cnt
}