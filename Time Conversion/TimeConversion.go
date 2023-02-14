package main

import(
	"fmt"
	"strconv"
	"strings"
)

func main(){
	str := "12:00:01PM"
	output := timeConversion(str)
	fmt.Println(output)
}

func timeConversion(s string) string {
    var outputtime string = ""
    if strings.Contains(s, "AM"){
        outputtime = ""
        amtime := strings.Split(s,"AM")
        slc := strings.Split(amtime[0],":")
        if slc[0]=="12"{
            outputtime = "00"
        } else {
            outputtime = slc[0]
        }
        outputtime = outputtime+":"+slc[1]+":"+slc[2]
    }

    if strings.Contains(s, "PM"){
        outputtime = ""
        pmtime := strings.Split(s,"PM")
        slc := strings.Split(pmtime[0],":")
        addhr := 0
        hrs,_ := strconv.Atoi(slc[0])

        if hrs == 12{
            addhr = 0
        } else {
            addhr = 12
        }
        hrs += addhr
        outputtime = strconv.Itoa(hrs)
        outputtime = outputtime + ":"+slc[1]+":"+slc[2]
    }
    return outputtime
}