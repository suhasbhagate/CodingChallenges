package main

import(
	"fmt"
)

func main(){
	nums := []int{2,1,2,1,2}
	k := 2
	res := subarraysWithKDistinct(nums,k)
	fmt.Println(res)
}


// func subarraysWithKDistinct(nums []int, k int) int {
//     solcnt:=0   
    
//     for i:=0; i<len(nums);i++{
// 		end := i
// 		m:= make(map[int]bool)
//         for end<len(nums) && len(m)<=k {
//             m[nums[end]]=true
//             end++
// 			if len(m)==k{
// 				solcnt++
// 			}            
//         }
//     }
//     return solcnt
// }

func subarraysWithKDistinct(nums []int, k int) int {
    solcnt:=0   
    m:= make(map[int]int)
	end := 0

    for i:= 0; i<=len(nums)-k; i++{
        for end<len(nums) && len(m)<=k {
            m[nums[end]]++
            end++
			if len(m)==k{
				//fmt.Println(m)
				solcnt++
			}
		}
		end--
		if len(m)>=k{			
			for len(m)>=k{
				if _,ok:= m[nums[end]];ok{
					m[nums[end]]--
					if m[nums[end]]==0{
						delete (m,nums[end])
					}
				}
				end--
			}
			end++
			m[nums[end]]++
			end++
		}

		if _,ok:= m[nums[i]];ok{
			m[nums[i]]--
			if m[nums[i]]==0{
				delete (m,nums[i])
			}
		}
		if len(m)==k{
			solcnt++
			//fmt.Println(m)
		}
    }
    return solcnt
}