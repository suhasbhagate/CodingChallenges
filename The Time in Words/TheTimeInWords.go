package main

import(
	"fmt"
)

func main(){
	var h int32 = 12
	var m int32 = 0
	fmt.Println(timeInWords(h,m))
}

func timeInWords(h int32, m int32) string {
    // Write your code here
    mp:= make(map[int32]string)
    mp[1] = "one"
    mp[2] = "two"
    mp[3] = "three"
    mp[4] = "four"
    mp[5] = "five"
    mp[6] = "six"
    mp[7] = "seven"
    mp[8] = "eight"
    mp[9] = "nine"
    mp[10] = "ten"
    mp[11] = "eleven"
    mp[12] = "twelve"
    mp[13] = "thirteen"
    mp[14] = "fourteen"
    mp[16] = "sixteen"
    mp[17] = "seventeen"
    mp[18] = "eighteen"
    mp[19] = "nineteen"
    mp[20] = "twenty"
    mp[21] = "twenty one"
    mp[22] = "twenty two"
    mp[23] = "twenty three"
    mp[24] = "twenty four"
    mp[25] = "twenty five"
    mp[26] = "twenty six"
    mp[27] = "twenty seven"
    mp[28] = "twenty eight"
    mp[29] = "twenty nine"
    
    var res string
    switch{
        case m == 0:
            res = mp[h]+" o' clock"
        case m ==1:
            res = mp[m]+" minute past "+mp[h]
        case m>1 && m<15:
            res = mp[m]+" minutes past "+mp[h]
        case m==15:
            res = "quarter past "+mp[h]        
        case m>15 && m<30:
            res = mp[m]+" minutes past "+mp[h]        
        case m==30:
            res = "half past "+mp[h] 
        case m>30 && m<45:
            res = mp[60-m]+" minutes to "+mp[h+1]
        case m==45:
            res = "quarter to "+mp[h+1]
        case m>45 && m<59:
            res = mp[60-m]+" minutes to "+mp[h+1]
        case m==59:
            res = mp[60-m]+" minute to "+mp[h+1]
    }
    return res
}