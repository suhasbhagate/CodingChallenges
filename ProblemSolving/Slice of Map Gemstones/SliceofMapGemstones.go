package main

import(
	"fmt"
)

func main(){
	arr := []string {"abcdde", "baccd", "eeabg'"}
	fmt.Println(gemstones(arr))
}


func gemstones(arr []string) int32 {
    l := len(arr)
    mp := make([]map[rune]int, l)
    for i, v:= range arr{
        mp[i] = make(map[rune]int)
        rns := []rune(v)
        for _,u := range rns{
            mp[i][u]++
        }
    }
    var cnt int32 = 0
    for w := range mp[0]{
        flg := 1
        for k:=1; k<len(mp);k++{
            if mp[k][w] == 0{
                flg = 0
                break 
            }
        }
        if flg == 1{
            cnt++
        }
    }
    return cnt
}