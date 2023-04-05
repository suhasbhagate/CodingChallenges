package main

import(
	"fmt"
)

func main(){
	hm:= Constructor()
	hm.Put(1,1)
	hm.Put(2,2)
	fmt.Println(hm.Get(1))
	fmt.Println(hm.Get(3))
	hm.Put(2,1)
	fmt.Println(hm.Get(2))
	hm.Remove(2)
	fmt.Println(hm.Get(2))
}

type Record struct{
	Key int
	Value int
}

type MyHashMap struct {

Data []Record
}


func Constructor() MyHashMap {
dt:= []Record{}
mhm:= MyHashMap{Data:dt}
return mhm
}


func (this *MyHashMap) Put(key int, value int)  {
flg:=false
for i:= 0; i<len(this.Data);i++{
	if this.Data[i].Key == key{
		flg=true
		this.Data[i].Value = value
	}
}
if flg == false{
	rc:=Record{Key:key, Value:value}
	this.Data = append(this.Data, rc)
}
}


func (this *MyHashMap) Get(key int) int {
for _,v:= range this.Data{
	if v.Key == key{
		return v.Value
	}
}
   return -1
}


func (this *MyHashMap) Remove(key int)  {
for i:=0; i< len(this.Data); i++{
	if this.Data[i].Key == key{
		this.Data = append(this.Data[:i],this.Data[i+1:]...)
	}
}
}


/**
* Your MyHashMap object will be instantiated and called as such:
* obj := Constructor();
* obj.Put(key,value);
* param_2 := obj.Get(key);
* obj.Remove(key);
*/