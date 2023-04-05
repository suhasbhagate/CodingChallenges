package main

import(
	"fmt"
)

func main(){
	hst:=Constructor()
	hst.Add(1)
	hst.Add(2)
	fmt.Println(hst.Contains(1))
	fmt.Println(hst.Contains(3))
	hst.Add(2)
	fmt.Println(hst.Contains(2))
	hst.Remove(2)
	fmt.Println(hst.Contains(2))
}


type MyHashSet struct {
    data []int
}


func Constructor() MyHashSet {
    dt:= []int{}
    mh:= MyHashSet{data:dt}
    return mh
}


func (this *MyHashSet) Add(key int)  {
    flg:=false
    for _,v:= range this.data{
        if v==key{
            flg = true
        }
    }
    if flg == false{
        this.data = append(this.data,key)
    }
}


func (this *MyHashSet) Remove(key int)  {
    for i:= 0; i<len(this.data); i++{
        if this.data[i]==key{
            this.data = append(this.data[:i],this.data[i+1:]...)
        }
    }
}


func (this *MyHashSet) Contains(key int) bool {
    for _,v:= range this.data{
        if v==key{
            return true
        }
    }
    return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */