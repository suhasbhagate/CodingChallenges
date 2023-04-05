package main

import(
	"fmt"
)

func main(){
	fmt.Println(isHappy(19))
}


func isHappy(n int) bool {
    sum := n
    m := make(map[int]bool)
    for {
        var slc []int
        num := sum
        for num>0{
            slc = append(slc, num%10)
            num/=10
        }
        sum = 0
        for _,v:= range slc{
            sum += v*v
        }
        if _,ok := m[sum]; ok{
            break
        }

        if sum==1{
            return true
        }
        m[sum] = true
    }
    return false
}