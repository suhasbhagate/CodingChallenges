package main

import(
	"fmt"
	"sort"
)

func main(){
	grid := []string{"abc", "ade", "efg"}
	res := gridChallenge(grid)
	fmt.Println(res)
}

func gridChallenge(grid []string) string {
    for i,v:= range grid{
        rn := []rune(v)
        var rnstr []string
        for _, u:= range rn{
            rnstr = append(rnstr,string(u))
        }
        sort.Strings(rnstr)
        var res string
        for _, w := range rnstr{
            res = res + w
        }
        grid[i] = res
    }
    
    rows := len(grid)
    columns := len(grid[0])
    var flg bool = true
    for i:=0; i<columns; i++{
        for j:= 0; j<rows-1 ;j ++{
            if grid[j][i] > grid [j+1][i]{
                flg = false
                return "NO"
            }
        }
    }
    if flg == true{
        return "YES"
    } else{
        return "NO"
    }
}