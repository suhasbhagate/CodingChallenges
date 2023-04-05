package main

import(
	"fmt"
)


func main(){
	num := []int{2,7,4}
	k := 181
	fmt.Println(addToArrayForm(num,k))
}


func addToArrayForm(num []int, k int) []int {
    maxlen := len(num)
    var num2 []int
    for k>0{
        num2 = append(num2, k%10)
        k = k/10
    }
    for i,j:=0,len(num)-1; i<j; i,j = i+1,j-1{
        num[i], num[j] = num[j],num[i]
    }
    if maxlen<len(num2){
        maxlen = len(num2)
    }
    if maxlen>len(num){
        for i:=len(num); i<maxlen;i++{
            num = append(num,0)
        }
    }
    if maxlen>len(num2){
        for i:=len(num2); i<maxlen;i++{
            num2 = append(num2,0)
        }
    }
    var res []int
    carry := 0
    for j:=0; j<maxlen;j++{
        sum := num[j]+num2[j]+carry
        if sum>9{
            sum = sum-10
            carry = 1
        } else{
            carry = 0
        }
        res = append(res,sum)
    }
    if carry>0{
        res = append(res,carry)
    }
    for i,j:=0,len(res)-1; i<j; i,j = i+1,j-1{
        res[i], res[j] = res[j],res[i]
    }
    return res
}