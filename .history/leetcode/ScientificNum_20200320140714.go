package leetcode

import(
	"strings"
	"fmt"
)

func ScientificToFloat(e string){
	var(
		sfloat	  float32
		snum	  int
	)	
	arr := strings.Split(e,"E")
	sfloat = float32(arr[0])
	snum = int(arr[1][1:])
	

}