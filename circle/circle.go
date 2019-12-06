package circle

import(
	"fmt"
)

type Circle struct{
	R 		int

}

type Rectangular struct{
	X		int
	Y		int
}

func PutCircle(){
	var re  Rectangular
	for i:=5;i>=0;i--{
		re.X = i
		re.Y = 5-i
		fmt.Println(re)
	}

}