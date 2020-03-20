package leetcode

import(
	"strings"
	"fmt"
)

func ScientificToFloat(e string){
	// var(
	// 	sfloat	  float32
	// 	snum	  int
	// )	
	arr := strings.Split(e,"E")

	for _,v := range arr{
		fmt.Println(v)
		}

}