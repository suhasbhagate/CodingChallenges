package main

import(
	"fmt"
	"strconv"
	"strings"
)


func main(){
	a := "0101"
	b := "101"
	fmt.Println(addBinary(a,b))

}


func addBinary(a string, b string) string {
	maxlen:= len(a)
	if len(b)>maxlen{
		maxlen = len(b)
	}
	var stra, strb []byte
	for i:=len(a)-1; i>=0; i--{
		stra = append(stra, a[i])
	}
	for j:=len(b)-1; j>=0; j--{
		strb = append(strb, b[j])
	}
	if len(a)<maxlen{
		for k:=0; k<maxlen-len(a);k++{
			stra = append(stra, '0')
		}
	}
	if len(b)<maxlen{
		for k:=0; k<maxlen-len(b);k++{
			strb = append(strb, '0')
		}
	}
	fmt.Println(string(stra))
	fmt.Println(string(strb))

	var sb strings.Builder
	ans:=0
	carry:=0
	for i:=0; i<maxlen;i++{
		if stra[i]=='0' && strb[i]=='0'{
			ans = carry
			sb.WriteString(strconv.Itoa(ans))
			carry = 0
			fmt.Println(rune(ans))
		}
		if stra[i]=='1' && strb[i]=='1'{
			if carry==1{
				ans = 1
				carry = 1
			}
			if carry==0{
				ans = 0
				carry = 1
			}
			sb.WriteString(strconv.Itoa(ans))
			fmt.Println(rune(ans))
		}
		if stra[i]!=strb[i]{
			if carry==1{
				ans = 0
				carry = 1
			}
			if carry==0{
				ans = 1
				carry = 0
			}
			sb.WriteString(strconv.Itoa(ans))
			fmt.Println(rune(ans))
		}
	}
    if carry==1{
        sb.WriteString(strconv.Itoa(carry))
    }

	fmt.Println(sb.String())
	rns := []rune(sb.String())
    for i,j:=0,len(rns)-1; i<j; i,j=i+1,j-1{
		rns[i],rns[j] = rns[j],rns[i]
	}
	return string(rns)
}