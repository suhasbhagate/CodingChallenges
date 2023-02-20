package main

import(
	"fmt"
	"math"
)

func main(){
	keyboards := []int32{3, 1}
	drives := []int32{5, 2, 8}
	var b int32 = 10
	fmt.Println(getMoneySpent(keyboards,drives,b))
}


func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
    /*
     * Write your code here.
     */
    if len(drives)==0 || len(keyboards)==0{
        return -1
    }
    if len(drives)==1 && len(keyboards)==1{
        if keyboards[0]+drives[0]>b{
            return -1
        } else{
            return keyboards[0]+drives[0]
        }
    } 
    kb := make([]int,len(keyboards))
    dr := make([]int,len(drives))
    
    var kblow int32= math.MaxInt32
    var drlow int32= math.MaxInt32
    
    for _,v:= range keyboards{
        if v<kblow{
            kblow = v
        }
    }
    
    for _,v:= range drives{
        if v<drlow{
            drlow = v
        }
    }
    if kblow+drlow >b{
        return -1
    }
     
      
    for i,v:= range keyboards{
        if v+drlow <= b{
            kb[i]=1
        } else{
            kb[i]=0
        }
    }
    
    for i,v:= range drives{
        if v+kblow <= b{
            dr[i]=1
        } else{
            dr[i]=0
        }
    }
    
    var max int32 = 0
    for i,u:= range keyboards{
        if kb[i]==1{
            for j,v:= range drives{
                if dr[j]==1{
                    if u+v <=b && u+v >max{
                        max = u+v
                    }
                }
            }
        }
    }
    return int32(max)
}