package main

import(
	"fmt"
)

func main(){
	grid:= []string{".......", "...O...", "....O..", ".......", "OO.....", "OO....."}
	var n int32 = 3
	fmt.Println(bomberMan(n,grid))
}

func bomberMan(n int32, grid []string) []string {
    // Write your code here
    gridslc:=make([][]int,len(grid))
    for i,u:= range grid{
        for _,v:= range u{
			var outint int
			if v=='O'{
				outint = 1
			} else{
				outint = 0
			}
            gridslc[i] = append(gridslc[i],outint)
        }
    }
	
	for m:=0; m<len(gridslc);m++{
		for n:= 0; n<len(gridslc[m]); n++{
			if gridslc[m][n]!=0{
				gridslc[m][n]++
			}
		}
	}

    var k int32
    for k=2; k<=n;k++{
        for m:=0; m<len(gridslc);m++{
			for n:= 0; n<len(gridslc[m]); n++{
				gridslc[m][n]++
			}
		}
		for i:= 0; i<len(gridslc);i++{
			for j:= 0; j<len(gridslc[i]); j++{
				if gridslc[i][j]==4{
					gridslc[i][j] = 0
					if i > 0 && gridslc[i-1][j]!=4{
						gridslc[i-1][j] = 0
					}
					if i < len(gridslc)-1 && gridslc[i+1][j]!=4{
						gridslc[i+1][j] = 0
					}
					if j > 0 && gridslc[i][j-1]!=4{
						gridslc[i][j-1] = 0
					}
					if j < len(gridslc[i])-1 && gridslc[i][j+1]!=4{
						gridslc[i][j+1] = 0
					}
				}
			}
		}
	}
	fmt.Println(gridslc)
	var resgrid []string
    for _,u:= range gridslc{
		var resstr []rune
        for _,v:= range u{
			if v==0 {
				resstr = append(resstr, '.')
			} else{
				resstr = append(resstr, 'O')
			}
        }
		resgrid = append(resgrid, string(resstr))
    }
	//fmt.Println(resgrid)
    return resgrid
}