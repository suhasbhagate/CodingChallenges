package main

import(
	"fmt"
)

func main(){
	digits := []int{9}
	fmt.Println(plusOne(digits))
}


func plusOne(digits []int) []int {
    i:=len(digits)-1
    digits[i] = digits[i]+1
    for j:=i; j>0;j--{
        if digits[j]>9{
            digits[j] = digits[j]-10
            digits[j-1]++
        }
    }
    var ans []int
    if digits[0]>9{
        ans = append(ans,digits[0]/10)
        ans = append(ans,digits[0]-10)
        for k:=1; k<len(digits);k++{
            ans = append(ans,digits[k])
        }
    } else{
        for k:=0; k<len(digits);k++{
            ans = append(ans,digits[k])
        }
    }
    return ans
}