package main

import(
	"fmt"
)

func main(){
	n:= 5
	Right(n)
	Left(n)
	Centre(n)
}

func Right(n int){
	for i:=1; i<=n;i++{
		for j:=1;j<=n-i;j++{
			fmt.Print(" ")
		}
		for k:=n-i+1; k<=n;k++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Left(n int){
	for i:=1; i<=n; i++{
		for j:=1; j<=i; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Centre(n int){
	for i:=1; i<=n;i++{
		for j:= 1; j<= n-i; j++{
			fmt.Print(" ")
		}
		for k:= 1; k<=i;k++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}