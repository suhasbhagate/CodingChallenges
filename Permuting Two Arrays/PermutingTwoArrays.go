package main

import(
	"fmt"
	"sort"
)

func main(){
	var k int32 = 1
	A := []int32{0, 1}
	B := []int32{0, 2}
	fmt.Println(twoArrays(k, A, B))	
}

func twoArrays(k int32, A []int32, B []int32) string {
    var slcA, slcB []int
    for _,v:= range A{
        slcA = append(slcA, int(v))
    }
    
    for _,v:= range B{
        slcB = append(slcB, int(v))
    }

    sort.Ints(slcA)
    sort.Sort(sort.Reverse(sort.IntSlice(slcB)))
    
    var ans string = "YES"
    for i:= range slcA{
        if slcA[i]+slcB[i]<int(k){
            ans = "NO"
			return ans
        }
    }
    return ans
}