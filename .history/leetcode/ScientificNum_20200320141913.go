package leetcode

import(
	"strings"
	"strconv"
	"fmt"
)

func ScientificToFloat(e string){
	var(
		sfloat	  float64
		snum	  int
		err		  error
	)	
	arr := strings.Split(e,"E")
	snum,err = strconv.Atoi(arr[0]); if err != nil{
		fmt.Println("输入非科学计数！")
	}
	sfloat,err = strconv.ParseFloat(arr[1][1:],64); if err != nil{
		fmt.Println("输入非科学计数！")
	}
	fmt.Println(snum)
	fmt.Println(sfloat)
	

}