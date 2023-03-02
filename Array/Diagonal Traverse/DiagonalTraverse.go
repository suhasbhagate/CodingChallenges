package main

import(
	"fmt"
)

func main(){
	mat:=[][]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(findDiagonalOrder(mat))
}


func findDiagonalOrder(mat [][]int) []int {
	m:= make(map[int][]int)
	for i:=0; i<len(mat);i++{
		for j:=0; j<len(mat[i]);j++{
			key := i+j
			m[key] = append(m[key], mat[i][j])
		}
	}
	//fmt.Println(m)
	var ans []int
	for k:=0; k<=len(mat)+len(mat[0])-2;k++{
		v:=m[k]
		if k%2 == 0{
			for l:=len(v)-1; l>=0; l--{				
				ans = append(ans, v[l])
			}
		}else{
			for l:=0; l<len(v);l++{
				ans = append(ans, v[l])
			}
		}
	}
	return ans
}