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
	snum = strconv.Itoa(arr[0])
	sfloat = strconv.ParseFloat(arr[1][1:],32)
	fmt.Println(snum)
	fmt.Println(sfloat)
	

}