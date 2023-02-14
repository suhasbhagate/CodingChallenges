package main

import(
	"fmt"
	"math"
)

func main(){
	a:= []int32{3, 2, 1, 2, 3}
	fmt.Println(minimumDistances(a))
}


func minimumDistances(a []int32) int32 {
    // Write your code here
    m:= make(map[int32]int32)
    for _,v:= range a{
        m[v]++
    }
    var slc []int32
    for k:= range m{
        if m[k]%2==0{
            slc = append(slc,k)
        }
    }
    if len(slc)==0{
        return -1
    }
    var min int32= math.MaxInt32
    for _,u:=range slc{
        i:=0
        for a[i]!=u{
            i++
        }
        st:= i
        i++
        for a[i]!=u{
            i++
        }
        end:=i
        if int(min)> end-st{
            min = int32(end-st)
        }
    }
    return min
}