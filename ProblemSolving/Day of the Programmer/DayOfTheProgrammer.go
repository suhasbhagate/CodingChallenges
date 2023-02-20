package main

import(
	"fmt"
	"strconv"
)

func main(){
	fmt.Println(dayOfProgrammer(2017))
}

func dayOfProgrammer(year int32) string {
    // Write your code here
    var res string
    if year>1918{
        if year%400==0 || (year%4==0 && year%100 !=0){
            res = "12"+".09."+strconv.Itoa(int(year))
        } else{
            res = "13"+".09."+strconv.Itoa(int(year))
        }
    }
    if year<1918{
        if year%4==0{
            res = "12"+".09."+strconv.Itoa(int(year))
        } else{
            res = "13"+".09."+strconv.Itoa(int(year))
        }
    }
    if year ==1918{
        res = "26"+".09."+strconv.Itoa(int(year))
    }
    return res
}