package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter String: ")
	str, _, err := reader.ReadLine()
	if err!=nil{
		fmt.Println("Wrong Input")
	}
	//str := "dvdf"
	n,res := LongestSubarray(string(str))
	fmt.Println("String: ",res)
	fmt.Println("Length: ",n)
}

func LongestSubarray(s string)(int,string){
	rns := []rune(s)
	maxlen := 0
	st := 0
	end := 0
	var res string
	if len(rns)<2{
		return len(rns),string(rns)
	} else{
		for i:= 0; i<len(rns);i++{
			flg := false
			indx := 0
			for j:=st; j<end;j++{
				if rns[j] == rns[i]{
					flg = true
					indx = j
					break
				}
			}

			if flg{
				if end-st > maxlen{
					maxlen = end-st
					res = string(rns[st:end])
				}
				st = indx+1
				end = i+1
			} else{
				end++
				if i == len(rns)-1{
					if end-st > maxlen{
						maxlen = end-st
						res = string(rns[st:end])
					}
					return maxlen, res
				}
			}
		}
		return maxlen,res
	}	
}
