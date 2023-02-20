package main

import(
	"fmt"
)

func main(){
	arr := []int32{5, 6, 8, 11}
	fmt.Println(balancedSums(arr))
}

func balancedSums(arr []int32) string {
    var ans string
    i := 0
    j := len(arr)-1
    var leftsum int32 = arr[i]
    var rightsum int32 = arr[j]
    for i<j{
        if leftsum < rightsum{
            i++
            leftsum += arr[i]
        } else{
            j--
            rightsum += arr[j]
        }
    }
    if leftsum == rightsum{
        ans = "YES"
    } else {
        ans = "NO"
    }
    return ans
}