package main

import(
	"fmt"
)

func main(){
	var d1,m1,y1 int32 = 14,7,2018
	var d2,m2,y2 int32 = 5,7,2018
	fmt.Println(libraryFine(d1,m1,y1,d2,m2,y2))
}

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
    // Write your code here
    var fine int32 = 0
    if y1-y2>0{
        fine += (y1-y2)* 10000
        return fine
    }
    if y1==y2 && m1-m2>0{
        fine += (m1-m2)* 500
        return fine
    }
    if y1==y2 && m1==m2 && d1-d2>0{
        fine += (d1-d2)*15
        return fine
    }
    return fine
}