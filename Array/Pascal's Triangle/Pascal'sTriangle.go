package main

import(
	"fmt"
)

func main(){
	fmt.Println(generate(5))
}


func generate(numRows int) [][]int {
    m:=make([][]int,numRows)
	for i:=0; i<numRows;i++{
		m[i] = make([]int, i+1)
			for j:=0; j<=i;j++{
				if j==0 || j==i{
					m[i][j]=1
				}else{
					m[i][j] = m[i-1][j-1]+m[i-1][j]
				}
			}
	}
	return m
}