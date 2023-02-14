package main

import(
	"fmt"
	"math"
)

func main(){
	var x int32 = 2
	var y int32 = 5
	var z int32 = 4
	fmt.Println(catAndMouse(x,y,z))	
}

func catAndMouse(x int32, y int32, z int32) string {
    var ans string
    if math.Abs(float64(x)-float64(z)) < math.Abs(float64(y)-float64(z)) {
        ans = "Cat A"
    } else if math.Abs(float64(x)-float64(z)) > math.Abs(float64(y)-float64(z)){
        ans = "Cat B"
    }else{
        ans = "Mouse C"
    }
    return ans
}