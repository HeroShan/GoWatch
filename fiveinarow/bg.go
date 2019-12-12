package fiveinarow

import(
	"fmt"
)
type Coordinat struct{
	x	int
	y	int
}

type Allinat struct{
	key		[]Coordinat
}



func AddCoordinat(coor Coordinat){
	var inat =  Allinat{}
	inat.key = append(inat.key,coor)
	 fmt.Println(inat)

}

