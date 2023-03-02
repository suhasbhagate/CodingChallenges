package main

import(
	"fmt"
)


func main(){
	matrix := [][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	fmt.Println(spiralOrder(matrix))
}


func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return nil
    }
	var ans []int
	ans=PrintSpiralOrder(matrix,0,0,len(matrix)-1, len(matrix[0])-1,ans)
	return ans
}


func PrintSpiralOrder(matrix [][]int, i, j, m, n int,ans []int)[]int{
	if i>m || j>n{
		return ans
	}
	for x:=j; x<=n;x++{
		ans = append(ans, matrix[i][x])
	}
	for y:=i+1; y<=m; y++{
		ans = append(ans, matrix[y][n])
	}
	if m!=i{
		for z:=n-1; z>=j; z--{
			ans = append(ans, matrix[m][z])
		}
	}
	if n!=j{
		for w:=m-1; w>i;w--{
			ans = append(ans, matrix[w][j])
		}
	}
	return PrintSpiralOrder(matrix,i+1,j+1,m-1,n-1,ans)
}