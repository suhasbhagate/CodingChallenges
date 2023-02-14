package main

import(
	"fmt"
)

func main(){
	var n int32 = 5
	staircaseRight(n)
	staircaseLeft(n)
	staircasePyramid(n)
}

func staircaseRight(n int32) {
    // Write your code here
    for i:= 1 ; i <= int(n); i++{
        for j:=1; j<=int(n)-i;j++{
            fmt.Print(" ")
        }
        for j:= int(n)-i+1; j<=int(n);j++{
            fmt.Print("#")
        }
        fmt.Println()
    }
}

func staircaseLeft(n int32){
	for i:=1; i<=int(n); i++{
		for j:=1; j<=i; j++{
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func staircasePyramid(n int32) {
    // Write your code here
    for i:= 1 ; i <= int(n); i++{
        for j:=1; j<=int(n)-i;j++{
            fmt.Print(" ")
        }
        for k:= 1; k<= i; k++{
            fmt.Print("# ")
        }
        fmt.Println()
    }
}