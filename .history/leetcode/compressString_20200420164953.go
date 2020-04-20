package leetcode

import(
	"strconv"
	"fmt"
)

func CompressString(S string) string {
    var(
        strlen      int
        tmp         string
        compress    string 
        strCount    int
    )
	strlen = len(S)
	fmt.Println(strlen)
    for i:=0; i<strlen; i++{
        if string(S[i]) == tmp{
            strCount++
        }else{
            if tmp != "" || i==(strlen-1){
				compress = compress+tmp+strconv.Itoa(strCount)
			}
			strCount = 1
            tmp = string(S[i])
		}
		fmt.Println(i)
	}
	return compress
    
}