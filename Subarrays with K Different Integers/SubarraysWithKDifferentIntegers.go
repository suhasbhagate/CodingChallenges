package main

import(
	"fmt"
)

func main(){
	nums := []int{1,2,1,2,3}
	k := 2
	res := subarrayWithKDistinct(nums,k)
	fmt.Println(res)
}


func subarrayWithKDistinct(nums []int, k int)int{
	solcnt := 0

	for i:=0; i<=len(nums)-k;i++{
		end := i
		m := make(map[int]bool)

		for len(m)<=k && end <len(nums){
			m[nums[end]] = true
			if len(m)==k{
				solcnt++
			}
			end++
		}	
	}
	return solcnt
}