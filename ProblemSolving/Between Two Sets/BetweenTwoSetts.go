package main

import(
	"fmt"
)

func main(){
	var a []int32 = []int32{2,6}
	var b []int32 = []int32{24,36}
	fmt.Println(getTotalX(a,b))
}
func getTotalX(a []int32, b []int32) int32 {
    // Write your code here
    max1 := a[0]
    min2 := b[0]
    for _, v:= range a{
        if max1<v{
            max1 = v
        }
    }
    for _, v:= range b{
        if min2>v{
            min2 = v
        }
    }
    cnt := 0
    for i:=max1; i<=min2;i+=max1{
        flg:= true
        for _,v:= range b{
            if v%i!=0{
                flg = false
            }
        }
        for _,v:= range a{
            if i%v != 0{
                flg = false
            }
        }
        if flg{
            cnt++
        }
    }
    return int32(cnt)
}