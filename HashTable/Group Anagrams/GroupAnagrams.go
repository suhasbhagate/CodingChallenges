package main

import(
	"fmt"
	"sort"
)

func main(){
	strs := []string{"eat","tea","tan","ate","nat","bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
    mp := make(map[string][]string)
    for _,s := range strs{
        rns := []rune(s)
        sort.Slice(rns, func(i,j int)bool{return rns[i]<rns[j]})
        mp[string(rns)] = append(mp[string(rns)], s)
    }
    var ans [][]string
    for k := range mp{
        ans = append(ans, mp[k])
    }
    return ans
}