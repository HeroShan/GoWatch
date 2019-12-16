package fiveinarow

import(
	 _"fmt"
	 "sort"
)
const(
	matrix = 50*50
	point  =  3
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

func inverted(tmp []int) bool {		//倒序检查
	var i int
	for k := len(tmp)-1; k>=0;k--{
		if k == 0{					//最后一个下标说明无法对比
			return false
		}
		if tmp[k]-1 == tmp[k]{		//说明值是连续数字
			i++
			if i == point{			//如果达到连续数就返回
				return true
			}
		}else{
			return false
		}
	}
	return false
}

func postive(tmp []int) bool {		//正序检查
	var i int
	for ck, cv := range tmp {
		if ck == len(tmp)-1{		//最后一个下标说明无法对比
			return false
		}
		if cv+1 == tmp[ck+1] {		//说明值是连续数字
			i++
			if i == point{			//如果达到连续数就返回
				return true
			}
		}else{
			return false
		}
	}
	return false
}

func lengthways(inat *Allinat,coor Coordinat) bool {
	var (
		max,min int = coor.x+4,coor.x-4
		tmp []int
	)
	if min < 0 {
		min = 0
	}
	for _,c := range inat.key{
		if (max >= c.x && c.y == coor.y) || (min >= c.x  && c.y == coor.y){
			tmp = append(tmp,c.x)
		}
	}
	sort.Ints(tmp)
	if (inverted(tmp) == true) || (postive(tmp) == true) {
		return true
	}
	return false
}

func crosswise(inat *Allinat,coor Coordinat) bool {
	var (
		max,min int = coor.y+4,coor.y-4
		tmp []int
	)
	for _,c := range inat.key{
		if (max >= c.y && c.x == coor.x) || (min >= c.y  && c.x == coor.x){
			tmp = append(tmp,c.y)
		}
	}
	sort.Ints(tmp)
	if (inverted(tmp) == true) || (postive(tmp) == true) {
		return true
	}
	return false
}

func IsFive(inat *Allinat,coor Coordinat) bool {
	ok := crosswise(inat,coor)
	ok2 := lengthways(inat,coor)
	//slope(inat)
	if ok == true || ok2 == true{
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

