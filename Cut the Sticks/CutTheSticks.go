package main

import(
	"fmt"
)

func main(){
	arr := []int32{5, 4, 4, 2, 2, 8}
	fmt.Println(cutTheSticks(arr))
}


func cutTheSticks(arr []int32) []int32 {
    // Write your code here
    var ans []int32
    ans = append(ans,int32(len(arr)))
    for{
        flg := true
        for _,v:= range arr{
            if v!= arr[0]{
                flg = false
            }
        }
        if flg{
            break
        }
        min:= arr[0]
        for _,v:= range arr{
            if min>v{
                min = v
            }
        }
        for i:= 0; i<len(arr);i++{
            arr[i] = arr[i] - min
        }

		var tmp []int32
        for i:=0 ; i<len(arr); i++{
            if arr[i]!=0{
                tmp = append(tmp, arr[i])
            }
        }
		arr = tmp
        fmt.Println(arr)
        ans = append(ans,int32(len(arr)))
    }
    return ans
}