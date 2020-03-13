package leetcode

import(
	"strings"
	"fmt"
)

func GcdOfStrings(str1 string, str2 string)  {
		tmpstr1 := strings.Split(str1,"")
		tmpstr2 := strings.Split(str2,"")
		fmt.Println(tmpstr1)
		var(
			vernier 		int
			vernierstr2		int
		)
		vernier = len(tmpstr1)-len(tmpstr2)
		for i:=0;i<=vernier;i++{
			vernierstr2 = len(tmpstr2)
			newstr := getCommonNum(tmpstr1,i,vernierstr2+i)
			// if newstr == str2{
			// 	return newstr
			// }
				fmt.Printf("new:%s \n",newstr)
		}


}

func getCommonNum(str []string, strIndex int, join int)string{
	if join == 1 {
		return 
	}
	strs:=str[strIndex:join]
	return strings.Join(strs,)
}