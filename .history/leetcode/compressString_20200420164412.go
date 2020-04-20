package leetcode

import(
	"strconv"
)

func CompressString(S string) string {
    var(
        strlen      int
        tmp         string
        compress    string 
        strCount    int
    )
    strlen = len(S)-1
    for i:=0; i<strlen; i++{
        if string(S[i]) == tmp{
            strCount++
        }else{
            if tmp != ""{
				compress = compress+tmp+strconv.Itoa(strCount)
			}
				
			
            
            strCount = 1
            tmp = string(S[i])
        }
	}
	return compress
    
}