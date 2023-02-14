package main

import(
	"fmt"
	"math"
)

func main(){
	//s := "haveaniceday"
	s:= "chillout"
	fmt.Println(encryption(s))
}

func encryption(s string)string{
    // Write your code here
	l := len(s)
    rows := int(math.Floor(math.Sqrt(float64(l))))
    cols := int(math.Ceil(math.Sqrt(float64(l))))
	if rows<cols && rows*cols < len(s){
		rows++
	}
	fmt.Println(l, rows, cols)
	var grid [][]byte
	i:=0
	for i<len(s){
		j:= i
		str:= make([]byte,cols)
		for j<len(s) && j<i+cols{
			ind:= j%cols
			str[ind] = s[j]
			j++
		}
		grid = append(grid, str)
		i = i+cols
	}
	fmt.Println(grid)
	res:=""
	for x:=0; x<cols;x++{
		var r []byte
		for y:=0;y<rows;y++{
			if grid[y][x]!=0{
				r = append(r, grid[y][x])
			}
		}
		fmt.Println(r)
		r = append(r, ' ')
		res = res+string(r)
	}
	res = res[:len(res)-1]
	return res
}