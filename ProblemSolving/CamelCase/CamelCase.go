package main

import(
	"fmt"
)

func main(){
	s:= "saveChangesInTheEditor"
	fmt.Println(camelcase(s))
}


func camelcase(s string) int32 {
    // Write your code here
    cnt:= 0
    for _,v:= range s{
        if v>='A' && v<='Z'{
            cnt++
        }
    }
    return int32(cnt+1)
}