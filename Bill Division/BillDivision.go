package main

import(
	"fmt"
)

func main(){
	var bill []int32 = []int32{3, 10, 2, 9}
	bonAppetit(bill,1,12)
}

func bonAppetit(bill []int32, k int32, b int32) {
    // Write your code here
    var sum int32 = 0
    var i int32
    for i= 0; i<int32(len(bill));i++{
        if i!= k{
            sum += bill[i]
        }
    }
    contri := sum/2
    if b == contri{
        fmt.Println("Bon Appetit")
    } else{
        fmt.Println(b-contri)
    }
}