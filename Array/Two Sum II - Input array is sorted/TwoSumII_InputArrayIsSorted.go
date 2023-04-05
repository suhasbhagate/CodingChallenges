package main

import(
	"fmt"
)


func main(){
	numbers := []int{2,7,11,15}
	fmt.Println(twoSum(numbers,9))
}


func twoSum(numbers []int, target int) []int {
    var ans []int
    for i:=0; i<len(numbers)-1; i++{
        for j:=i+1; j<len(numbers);j++{
            if numbers[i]+numbers[j] == target{
                ans = append(ans, i+1,j+1)
                return ans
            }
        }
    }
    return ans
}