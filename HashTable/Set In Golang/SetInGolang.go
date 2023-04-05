package main

import(
	"fmt"
)

func main(){
	st := CreateSet()
	st.Add(5)
	fmt.Println(st)
	st.Add(7)
	fmt.Println(st)
	st.Add("ABC")
	fmt.Println(st)
	fmt.Println(st.Contains(8))
	if st.Contains(5){
		st.Remove(5)
	}
	fmt.Println(st)
}

type Set struct{
	mp map[interface{}]struct{}
}

func CreateSet() Set{
	m := make(map[interface{}]struct{})
	s:= Set{mp:m}
	return s
}

func (s Set) Add(data interface{}){
	s.mp[data] = struct{}{}
}

func (s Set) Remove(data interface{}){
	// for k:= range s.mp{
	// 	if k == data{
	// 		delete(s.mp,k)
	// 	}
	// }
	if _,ok:=s.mp[data];ok{
		delete(s.mp,data)
	}
}

func (s Set) Contains (data interface{})bool{
	// for k:= range s.mp{
	// 	if k==data{
	// 		return true
	// 	}
	// }
	if _,ok:=s.mp[data];ok{
		return true
	}
	return false
}
