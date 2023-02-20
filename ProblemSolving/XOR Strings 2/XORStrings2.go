package main

import(
	"fmt"
)

func main(){
	s := "10101"
	t := "00101"
	res := strings_xor(s,t)
	fmt.Println(res)
}

func strings_xor(s, t string)string {
    res := "";
    for i := 0; i < len(s); i++ {
        if s[i] == t[i]{
            res = res+"0";
		} else{
            res = res+"1";
		}
    }
    return res;
}