package main

import(
	"fmt"
)

func main(){
	strs := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(strs))
}


func longestCommonPrefix(strs []string) string {
    if len(strs)==0{
        return ""
    }
    
    var common string
    
    for i:=0; i<len(strs[0]);i++{
        match := true
        
        for j:=1; j<len(strs);j++{
            if len(strs[j])-1<i{
                match = false
                break
            }
            if strs[0][i]!=strs[j][i]{
                match = false
                break
            }

        }
        if match ==true{
            common+=string(strs[0][i])
        } else{
            break
        }
    }
    return common
}