package leetcode

import(
	"strings"
	"math"
	"fmt"
)

func GcdOfStrings(str1 string, str2 string) string {
		tmpstr1 := strings.Split(str1,"")
		tmpstr2 := strings.Split(str2,"")
		var(
			vernier 		int
			vernierstr2		int
			diff			int
		)
		vernier = len(tmpstr1)-len(tmpstr2)
		vernierstr2 = len(tmpstr2)
		
		newstr2,str2len := getVicestr(tmpstr2,vernierstr2)
		if str2len != 0{
			str2 		= newstr2
			vernierstr2 = str2len
		}
		fmt.Printf("newstr2:%s ,%d\n",newstr2,str2len)
		for i:=0;i<=vernier;i++{
			getnewstr  := getCommonNum(tmpstr1,i,vernierstr2+i)
			
			fmt.Printf("newstr:%s ,str2:%s,%d,%d\n",getnewstr,str2,vernierstr2,i)
			if vernierstr2 == 1{
				break
			}else{
				if getnewstr == str2{
					return getnewstr
				}else if i==vernier {
					if diff == vernierstr2-1{
						break
					}
					 vernierstr2--
					// diff ++
					// str2 = getCommonNum(tmpstr2,0,vernierstr2-diff)
					i = -1
				}else if i==vernier {
					break
				}
			}
			
		}
		return ""

}

func getVicestr(str2 []string ,len int) (string,int){
	var vernierstr2 int
	tmp := math.Floor(float64(len/2))
	vernierstr2 = int(tmp)
	for i:=0;i<=len;i++{
	   if vernierstr2 == 1{
		   break
	   }
	   newstr1 := getCommonNum(str2,i,vernierstr2)
	   newstr2 := getCommonNum(str2,i+vernierstr2,i+vernierstr2+vernierstr2) 
	   if newstr1 == newstr2 {
		   return newstr1,vernierstr2
	   }else{
		   i = 0
		   vernierstr2--
	   }
	}
	return "",0
}

func getCommonNum(str []string, strIndex int, join int)string{
	strs:=str[strIndex:join]
	return strings.Join(strs,"")
}

