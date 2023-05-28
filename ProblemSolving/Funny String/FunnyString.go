package main

import (
	"fmt"
	"math"
)

func main(){
	s := "abx"
	fmt.Println(funnyString(s))
}


func funnyString(s string) string {
    srns := []rune(s)
    trns := make([]rune,len(srns))
    copy(trns,srns)
    for i, j:=0,len(srns)-1;i<j;i,j = i+1,j-1{
        trns[i],trns[j] = trns[j],trns[i]
    } 
    var sdiff, tdiff []int
    for j:=0; j<len(srns)-1;j++{
		sdiff = append(sdiff,int(math.Abs(float64(srns[j+1]-srns[j]))))
    }
    for k:=0; k<len(trns)-1;k++{
		tdiff = append(tdiff,int(math.Abs(float64(trns[k+1]-trns[k]))))
    }
    flg := true
    for i:=0; i<len(sdiff);i++{
        if sdiff[i]!= tdiff[i]{
            flg = false
        }
    }
    if flg==true{
        return "Funny"
    } else{
        return "Not Funny"
    }
}