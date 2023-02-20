package main

import(
	"fmt"
	"math"
)

func main(){
	var arr [][]int32 = [][]int32{{11, 2, 4},{4, 5, 6}, {10, 8, -12}}
	fmt.Println(diagonalDifference(arr))
}

func diagonalDifference(arr [][]int32) int32 {
    // Write your code here
    var leftsum int32 = 0
    var rightsum int32 = 0
    for i:=0; i<len(arr); i++{
        for j:=0; j<len(arr[0]); j++{
            if i==j{
                leftsum+=arr[i][j]
            }
            if i+j == len(arr)-1{
                rightsum+=arr[i][j]
            }
        }
    }
    return int32(math.Abs(float64(leftsum)-float64(rightsum)))
}