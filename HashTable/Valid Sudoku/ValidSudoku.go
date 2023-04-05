package main

import(
	"fmt"
)

func main(){
	board := [][]byte{{'5','3','.','.','7','.','.','.','.'}, {'6','.','.','1','9','5','.','.','.'},{'.','9','8','.','.','.','.','6','.'},{'8','.','.','.','6','.','.','.','3'},{'4','.','.','8','.','3','.','.','1'},{'7','.','.','.','2','.','.','.','6'},{'.','6','.','.','.','.','2','8','.'},{'.','.','.','4','1','9','.','.','5'},{'.','.','.','.','8','.','.','7','9'}}
	fmt.Println(isValidSudoku(board))
}


func isValidSudoku(board [][]byte) bool {
    for _, row := range board{
        mp := make(map[byte]int)
        for _,val := range row{
            if val != '.'{
                if _,ok := mp[val];ok{
                    return false
                } else{
                    mp[val] = 1
                }
            }
        }
    }
    
    for i := range board{
        mp := make(map[byte]int)
        for j:= range board[i]{
            val:= board[j][i]
            if val != '.'{
                if _,ok := mp[val];ok{
                    return false
                } else{
                    mp[val] = 1
                }
            }
        }
    }
    
    for i:=0; i<3; i++{
        for j:=0; j<3; j++{
            mp := make(map[byte]int)
            for k:= i*3; k<(i+1)*3; k++{
                for l:= j*3; l<(j+1)*3; l++{
                    val:= board[k][l]
                    if val != '.'{
                        if _,ok := mp[val];ok{
                            return false
                        } else{
                            mp[val] = 1
                        }
                    }
                }
            }                 
        }
    }
    return true
}