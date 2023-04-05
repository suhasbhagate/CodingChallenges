package main

import(
	"fmt"
	"math"
)

func main(){
	list1 := []string{"happy","sad","good"}
	list2 := []string{"sad","happy","good"}
	fmt.Println(findRestaurant(list1, list2))
}


func findRestaurant(list1 []string, list2 []string) []string {
    m1 := make(map[string]int)
    m2 := make(map[string]int)
    for i,v := range list1{
        m1[v] = i
    }
    for i,v := range list2{
        m2[v] = i
    }
    
    mp := make(map[string]int)
    for k,v1 := range m1{
        if v2,ok := m2[k]; ok{
            mp[k] = v1+v2
        }
    }
    
    min := math.MaxInt
    for k:= range mp{
        if mp[k]<min{
            min = mp[k]
        }
    }
       
    var ans []string
    for k:= range mp{
        if mp[k] == min{
            ans = append(ans, k)
        }
    }
    return ans    
}