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

func crosswise(inat *Allinat) bool {
	sucfice := make([][]int,100)
	for _,coslice := range inat.key{
		sucfice[coslice.x] = append(sucfice[coslice.x],coslice.y)
		if len(sucfice[coslice.x]) == 5{
			fmt.Printf("%T",sucfice[coslice.x])
			return true
		}
	}
	return false
}

func IsFive(inat *Allinat){
	ok := crosswise(inat)
	fmt.Println(ok)
}

func (inat *Allinat)AddCoordinat(coor Coordinat){
	for _,coslice := range inat.key{
		if coslice == coor {
			return 
		}
	}
	inat.key = append(inat.key,coor)
	IsFive(inat)
}

