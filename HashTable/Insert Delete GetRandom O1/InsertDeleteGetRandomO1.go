package main

import(
	"fmt"
)

func main(){
	rn := Constructor()
	fmt.Println(rn.Insert(1))
	fmt.Println(rn.Insert(1))
	fmt.Println(rn.Insert(2))
	fmt.Println(rn.Remove(2))
	fmt.Println(rn.Remove(2))
	fmt.Println(rn.GetRandom())
}


type RandomizedSet struct {
    mp map[int]bool
}


func Constructor() RandomizedSet {
    dt := make(map[int]bool)
    rs := RandomizedSet{mp: dt}
    return rs
}


func (this *RandomizedSet) Insert(val int) bool {
    if _,ok:= this.mp[val];ok{
        return false
    } else{
        this.mp[val] = true
        return true
    }
}


func (this *RandomizedSet) Remove(val int) bool {
    if _,ok := this.mp[val];ok{
        delete (this.mp,val)
        return true
    } else{
        return false
    }
}


func (this *RandomizedSet) GetRandom() int {
    randomkey := 0
    for k:= range this.mp{
        randomkey = k
        break
    }
    return randomkey
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */