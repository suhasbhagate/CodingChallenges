package main

import(
	"fmt"
)

func main(){
	s := "aaab"
	fmt.Println((palindromeIndex(s)))
}


func palindromeIndex(s string) int32 {
    rns := []rune(s)
    var i int32= 0
    var j int32= int32(len(s))-1
    cnt := 0
    var ans int32 = -1
    for i<j {        
        if rns[i]!=rns[j]{
            if cnt>0{
                ans = -1
                break
            } else{
                cnt++
                if rns[i]==rns[j-1] && rns[i+1]!=rns[j]{
                    ans = j
                    j--
                } else if rns[i]==rns[j-1] && rns[i+1]==rns[j-2]{
                    ans = j
                    j--
                } else if rns[i]==rns[j-1] && rns[i+1]==rns[j]{
                    ans = i
                    i++
                } else{
                    ans = i
                    i++
                }
            }
        }
        i++
        j--
    }
    return ans
}