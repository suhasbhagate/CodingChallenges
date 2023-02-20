package main

import(
	"fmt"
	"math"
)

func main(){
	gifts := []int{25,64,9,4,100}
	k := 4
	fmt.Println(pickGifts(gifts,k))
}

func pickGifts(gifts []int, k int) int64 {
    for i:=1; i<=k;i++{
        max:= 0
        indx := 0
        for j,v:= range gifts{
            if v>max{
                max = v
                indx = j
            }
        }
        gifts[indx] = int(math.Sqrt(float64(max)))
    }
    var sum int64 = 0
    for _,v:= range gifts{
        sum += int64(v)                              
    }
    return sum
}