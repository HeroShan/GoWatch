package leetcode

import(
    "strconv"
)

func compressString(S string) string {
    var(
        strlen      int,
        tmp         string,
        compress    string,
        strCount    int
    )
     := len(S)-1
    var tmp string 
    for i:=0; i<strlen; i++{
        if S[i]==tmp{
            strCount++
        }else{
            
            compress = compress+tmp+strconv.Atoi(strCount)
            strCount = 1
            tmp = S[i]
        }
       S[i] = tmp 
    }
    
}