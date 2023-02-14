package main

import(
	"fmt"
)

func main(){
	c:= []int32{0,0,1,0}
	var k int32 = 2
	fmt.Println(jumpingOnClouds(c,k))
}


func jumpingOnClouds(c []int32, k int32) int32 {
    var energy int32 = 100
    var dst int = 0
    for {
        dst = dst+int(k)
        if dst>=len(c){
            dst = dst - len(c)
        }
        energy = energy -1
        if c[dst]==1{
            energy = energy-2
        }
        if dst == 0{
            break
        }        
    }
    return energy
}