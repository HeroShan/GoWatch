package fiveinarow

import(
	"fmt"
)
const(
	matrix = 50*50
	point  =  5
)
type Coordinat struct{
	x	int
	y	int
}

type Allinat struct{
	key		[]Coordinat
}

func slope(inat *Allinat){
	// sucfice := make([][]int,matrix)
	// tmp := Coordinat{}
	//var sulen int
	for k,coslice := range inat.key{
		fmt.Println(k,coslice)
	}
	//fmt.Println(cap(sucfice))
}

func lengthways(inat *Allinat) bool {
	sucfice := make([][]int,matrix)
	for _,coslice := range inat.key{
		sucfice[coslice.y] = append(sucfice[coslice.y],coslice.x)
		if len(sucfice[coslice.y]) == point{
			return true
		}
	}
	return false
}

func crosswise(inat *Allinat) bool {
	sucfice := make([][]int,matrix)
	for _,coslice := range inat.key{
		sucfice[coslice.x] = append(sucfice[coslice.x],coslice.y)
		if len(sucfice[coslice.x]) == point{
			return true
		}
	}
	return false
}

func IsFive(inat *Allinat){
	ok := crosswise(inat)
	ok2 := lengthways(inat)
	slope(inat)
	if ok == true || ok2 == true{
		return 
	}
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
