package main

import(
	"fmt"
)

func main(){
	matrix := [][]int32{{112, 42, 83, 119}, {56, 125, 56, 49},{15, 78, 101, 43}, {62, 98, 114, 108}}
	var res int32
	res = flippingMatrix(matrix)
	fmt.Println(res)
}


func flippingMatrix(matrix [][]int32) int32{
	end := len(matrix)-1
	mid := len(matrix)/2
	var sum int32	
	for i:= 0; i<mid;i++{
		for j:=0; j<mid; j++{
			slc := make([]int32,4)
			slc = append(slc,matrix[i][j],matrix[i][end-j],matrix[end-i][j],matrix[end-i][end-j])
			var max int32
			for _, v := range slc{
				if v>max{
					max = v
				}
			}
			sum += max
		}
	}
	return sum
}