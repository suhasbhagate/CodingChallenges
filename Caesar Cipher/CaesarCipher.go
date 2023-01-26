package main

import(
	"fmt"
)

func main(){
	str :="www.abc.xy"
	var k int32 = 87
	res := caesarCipher(str,k)
	fmt.Println(res)
}

func caesarCipher(s string, k int32) string {
	k = k%26
	rns := []rune(s)
	var res []rune
	for _, v := range rns{
		var ans rune
		if v >= 'a' && v <= 'z'{
			ans = v + k
			if ans>'z'{
				ans -=26
			}
			res = append(res, ans)
		} else if v >= 'A' && v <= 'Z'{
			ans = v + k
			if ans>'Z'{
				ans -=26
			}
			res = append(res, ans)
		} else{
			res = append(res, v)
		}		
	}    
	return string(res)
}