package main

import (
	"fmt"
	"strconv"
	"strings"
)

var month = []int{31,28,31,30,31,30,31,31,30,31,30,31}
func main(){
	u := "02-08-2019" 
	v := "02-08-2020"
	uslc := strings.Split(u, "-")
	vslc := strings.Split(v, "-")
	d1,_ := strconv.Atoi(uslc[0])
	m1,_ := strconv.Atoi(uslc[1])
	y1,_ := strconv.Atoi(uslc[2])
	d2,_ := strconv.Atoi(vslc[0])
	m2,_ := strconv.Atoi(vslc[1])
	y2,_ := strconv.Atoi(vslc[2])
	res := findPrimeDates(d1,m1,y1,d2,m2,y2)
	fmt.Println(res)
}

func updateLeapYear(year int){
	if year % 4 == 0{
		month[1] = 29
	}
}

func findPrimeDates(d1, m1, y1, d2, m2, y2 int)int{	
	res := 0

	for{
		x := d1
		x = x * 100 + m1
		x = x * 10000 + y1

		if x % 4 == 0 || x % 7 ==0{
			res +=1
		}

		if(d1 == d2 && m1 == m2 && y1==y2){
			break
		}

		updateLeapYear(y1)
		
		d1 = d1+1
		if d1>month[m1-1]{
			m1 +=1
			d1 = 1
		}
		if m1>12{
			y1 +=1
			m1 = 1
		}
	}
	return res
}