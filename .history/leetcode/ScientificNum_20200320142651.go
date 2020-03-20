package leetcode

import(
	"strings"
	"strconv"
	"fmt"
)

func ScientificToFloat(e string) {
	var(
		// sfloat	  float64
		// snum	  int
		err		  error
	)	
	arr := strings.Split(e,"E")
	snum,err = strconv.Atoi(arr[1]); if err != nil{
		fmt.Println("Num输入非科学计数！")
	}
	sfloat,err = strconv.ParseFloat(arr[0][0:],64); if err != nil{
		fmt.Println("Float输入非科学计数！")
	}
	fmt.Printf("1%s   2%s",arr[0],arr[1])
	

}