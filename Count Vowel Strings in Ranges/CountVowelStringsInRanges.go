package main

import(
	"fmt"
)

func main(){
	words := []string{"aba","bcb","ece","aa","e"}
	queries := [][]int{{0,2},{1,4},{1,1}}
	fmt.Println(vowelStrings(words, queries))
}


func IsVowel(b byte)bool{
    if b=='a' || b=='e' || b=='i' || b=='o' || b=='u'{
        return true
    } else{
        return false
    }
}

func vowelStrings(words []string, queries [][]int) []int {
    slc := make([]int, len(words))
    for i,v:=range words{
        if IsVowel(byte(v[0])) && IsVowel(byte(v[len(v)-1])){
            slc[i] = 1
        }
    }
    
    ans:= make([]int, len(queries))
    for y,u:= range queries{
        sum := 0
        for x:= u[0]; x<=u[1];x++{
            sum += slc[x]
        }
        ans[y] = sum
    }
    return ans
}