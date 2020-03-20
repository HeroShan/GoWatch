package leetcode

import(
	"strings"
	"strconv"
	"fmt"
	"math"
)

func ScientificToFloat(e string) ( string){
	var(
		sfloat	  float64
		snum	  int
		err		  error
	)	
	arr := strings.Split(e,"E")
	snum,err = strconv.Atoi(arr[1]); if err != nil{
		fmt.Println("Num输入非科学计数！")
	}
	sfloat,err = strconv.ParseFloat(arr[0],64); if err != nil{
		fmt.Println("Float输入非科学计数！")
	}
	sfloat = sfloat * math.Pow10(snum)
	return string(sfloat)
	
}