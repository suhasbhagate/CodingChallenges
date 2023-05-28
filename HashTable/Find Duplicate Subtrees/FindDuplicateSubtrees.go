package main

import(
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}

func main(){
	var root *TreeNode = &TreeNode{Val:2,Left:&TreeNode{1,nil,nil},Right:&TreeNode{1,nil,nil}}
	rr := findDuplicateSubtrees(root)
	for _,v:= range rr{
		fmt.Println(v.Val)
	}
}


func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    var res []*TreeNode
    mp := make(map[string][]*TreeNode)
    _ = Traverse(root, mp)
    //fmt.Println(res)
    for _, value := range mp {
        if len(value) > 1 {
            res = append(res, value[0])
        }
    }    
    return res
}

func Traverse(root *TreeNode, mp map[string][]*TreeNode) string{
    if root==nil{
        return ""
    }
    
    key := strconv.Itoa(root.Val)+"#"+Traverse(root.Left, mp)+"#"+Traverse(root.Right,mp)
    //fmt.Println(key)

    mp[key] = append(mp[key], root)
    
    return key    
}