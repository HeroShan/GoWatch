package leetcode

import(
	"strings"
	"strconv"
	"fmt"
)

func ScientificToFloat(e string){
	var(
		sfloat	  float32
		snum	  int
	)	
	arr := strings.Split(e,"E")
	snum,_ = strconv.Atoi(arr[0])
	sfloat,_ = strconv.ParseFloat(arr[1][1:],32)
	fmt.Println(snum)
	fmt.Println(sfloat)
	

}