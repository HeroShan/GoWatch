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



func (inat *Allinat)AddCoordinat(coor Coordinat){
	for _,coslice := range inat.key{
		if coslice == coor {
			fmt.Println(coslice)
			return 
		}
	}
	inat.key = append(inat.key,coor)

}

