package leetcode

import(
	"strings"
	"math"
)

func GcdOfStrings(str1 string, str2 string) string {
		tmpstr1 := strings.Split(str1,"")
		tmpstr2 := strings.Split(str2,"")
		var(
			vernier 		int
			vernierstr2		int
		)
		vernier = len(tmpstr1)-len(tmpstr2)
		vernierstr2 = len(tmpstr2)


		for i:=0;i<=vernier;i++{
			newstr := getCommonNum(tmpstr1,i,vernierstr2+i)
			if vernierstr2 == 1{
				break
			}else{
				if newstr == str2{
					return newstr
				}else if i==vernier {
					vernierstr2--
					i = 0
				}else if i==vernier {
					break
				}
			}
		}
		return ""

}

func getCommonNum(str []string, strIndex int, join int)string{
	strs:=str[strIndex:join]
	return strings.Join(strs,"")
}

func getVicestr(str2 []string ,len int){
	 var vernierstr2 int
	 tmp := math.Floor(float64(len/2))
	 vernierstr2 = int(tmp)
	 for i:=0;i<=len;i++{
		getCommonNum(str2,i,vernierstr2) == getCommonNum(str2,i+vernierstr2,vernierstr2)
	 }

}