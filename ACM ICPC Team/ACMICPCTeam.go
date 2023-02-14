package main

import (
	"fmt"
	"math"
	//"strconv"
)

func main(){
	topics := []string{"10101", "11110", "00010"}
	//topics := []string{"10101", "11100", "11010", "00101"}
	fmt.Println(acmTeam(topics))
	// //acmTeam(topics)
	// fmt.Println(GetORResult("10101","11110"))
	// fmt.Println(GetBitsSet("10101"))
}

func GetORResult(a, b string) string{
	lena := len(a)
	lenb := len(b)
	maxlen := int(math.Max(float64(lena), float64(lenb)))
	var res string
	for i:= 0; i< maxlen; i++{
		if a[i]=='1' || b[i]=='1'{
			res = res+"1"
		} else{
			res = res +"0"
		}
	}
	return res
}

func GetBitsSet(s string)int32{
	var cnt int32= 0
	for _,v:= range s{
		if v=='1'{
			cnt++
		}
	}
	return cnt
}

func acmTeam(topic []string) []int32 {
    // Write your code here
	var orresultstring []string
	var orresultbitcnt []int32
	for i:=0; i<len(topic);i++{
		for j:=i+1;j<len(topic);j++{
			orres:= GetORResult(topic[i],topic[j])
			orcnt:= GetBitsSet(orres)
			orresultstring = append(orresultstring, orres)
			orresultbitcnt = append(orresultbitcnt, orcnt)
		}
	}
	var max int32 = 0
	for _,v:= range orresultbitcnt{
		if v>max{
			max = v
		}
	}
	var maxcnt int32 = 0
	for _,v:= range orresultbitcnt{
		if v==max{
			maxcnt++
		}
	}
	return []int32{max,maxcnt}
}

// func acmTeam(topic []string) []int32 {
//     // Write your code here
//     var numslc []int64
//     for _,v:= range topic{
//         var num int64
//         num,_ = strconv.ParseInt(v, 2, 64)
//         numslc = append(numslc,num)
//     }
// 	fmt.Println(numslc)
// 	var orresult []int64
// 	for i:= 0 ; i<len(numslc);i++{
// 		for j:=i+1; j<len(numslc);j++{
// 			orr := numslc[i]|numslc[j]
// 			orresult = append(orresult, orr)
// 		}
// 	}
// 	var max int64 = 0
// 	for _,v:= range orresult{
// 		if v > max{
// 			max = v
// 		}
// 	} 
// 	var cnt int32 = 0
// 	for _,v:= range orresult{
// 		if v == max{
// 			cnt++
// 		}
// 	}
// 	var bitcnt int32 =0
// 	for max!=0{
// 		if max & 1>0{
// 			bitcnt++
// 		}
// 		max >>=1
// 	}
// 	return []int32{bitcnt,cnt}
// }