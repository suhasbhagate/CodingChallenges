package main

import(
	"fmt"
)

func main(){
	h := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	word := "abc"
	fmt.Println(designerPdfViewer(h,word))
}

func designerPdfViewer(h []int32, word string) int32 {
    // Write your code here
    bslc := []byte(word)
    var max int32 = 0
    for _, v:= range bslc{
        if h[v-97]>max{
            max = h[v-97]
        }
    }
    return max * int32(len(word))
}