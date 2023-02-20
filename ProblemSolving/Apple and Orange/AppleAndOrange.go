package main
import(
	"fmt"
)

func main(){
	var apples []int32 = []int32{2, 3, -4}
	var oranges []int32 = []int32{3, -2, -4}
	countApplesAndOranges(7, 10, 4, 12, apples,oranges)
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
    // Write your code here
    for i:=0; i<len(apples); i++{
        apples[i] = apples[i]+a
    }
    for i:=0; i<len(oranges); i++{
        oranges[i] = oranges[i]+b
    }
    cntapple:= 0
    cntorange:= 0
    for _, v:= range apples{
        if v>=s && v<=t{
            cntapple++
        } 
    }
    for _, v:= range oranges{
        if v>=s && v<=t{
            cntorange++
        } 
    }
    fmt.Println(cntapple)
    fmt.Println(cntorange)
}