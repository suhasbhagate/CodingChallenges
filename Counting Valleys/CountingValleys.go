package main

import(
	"fmt"
)

func main(){
	var steps int32 = 8
	path := "DDUUUUDD"
	fmt.Println(countingValleys(steps,path))
}

func countingValleys(steps int32, path string) int32 {
    // Write your code here
    var slc []int32
    slc = append(slc,0)
    var val int32= 0
    var cntval int32 = 0
    for _,v:= range path{
        if v=='U'{
            val++
        } else if v=='D'{
            val--
        }
        slc = append(slc,val)
    }
    slc = append(slc,0)
    st := 0
    end := 0
    for i:=0; i<len(slc)-2;i++{
        if slc[i]==0{
            st = i
            end = i+1
            for end<len(slc) && slc[end]!=0{
                end++
            }
            flg := true
            for m:=st+1;m<end;m++{
                if slc[m]>0{
                    flg = false
                }
            }
            if flg{
                cntval++
            }
        }
    }
    return cntval
}