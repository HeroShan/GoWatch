package fiveinarow

import(
	 "fmt"
	 "sort"
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

// func slope(inat *Allinat){
// 	sucfice := make([][]int,matrix)
// 	tmp := Coordinat{}
// 	var sulen int
// 	for _,coslice := range inat.key{
		
// 	}
// 	//fmt.Println(cap(sucfice))
// }

func InArray(need Coordinat, needArr []Coordinat) bool {
	for _,v := range needArr{
	   if need == v{
		   return true
	   }
   }
   return false
}

func lengthways(inat *Allinat,coor Coordinat) bool {
	var (
		max,min int = coor.x+4,coor.x-4
		i int
		tmp []int
	)
	for _,c := range inat.key{
		if (max > c.x && c.y == coor.y) || (min > c.x  && c.y == coor.y){
			tmp = append(tmp,c.x)
		}
	}
	sort.Ints(tmp)
	for ck, cv := range tmp {
		if ck == len(tmp)-1{
			return false
		}
		if cv+1 == tmp[ck+1] {
			i++
		}
		if i == point{
			fmt.Println("X :",coor)
			return true
		}
	}
	return false
}

func crosswise(inat *Allinat,coor Coordinat) bool {
	var (
		max,min int = coor.y+4,coor.y-4
		i int
		tmp []int
	)
	for _,c := range inat.key{
		if (max > c.y && c.x == coor.x) || (min > c.y  && c.x == coor.x){
			tmp = append(tmp,c.y)
		}
	}
	sort.Ints(tmp)
	for ck, cv := range tmp {
		if ck == len(tmp)-1{
			return false
		}
		if cv+1 == tmp[ck+1] {
			i++
		}
		if i == point{
			fmt.Println("Y :",coor)
			return true
		}
	}
	return false
}

func IsFive(inat *Allinat,coor Coordinat) bool {
	ok := crosswise(inat,coor)
	ok2 := lengthways(inat,coor)
	//slope(inat)
	if ok == true && ok2 == true{
		return true
	}
	return false
}

func (inat *Allinat)AddCoordinat(coor Coordinat){
	for _,coslice := range inat.key{
		if coslice == coor {
			return 
		}
	}
	c := IsFive(inat,coor)
	if c == false{		//not finish five
		inat.key = append(inat.key,coor)
		return 
	}
	return 
	
}

