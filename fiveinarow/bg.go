package fiveinarow

import(
	 "fmt"
	 "sort"
)
const(
	matrix = 100*100
	point  =  2
)
type Coordinat struct{
	x	int
	y	int
}

type Allinat struct{
	key		[]Coordinat
}

func InArray(need Coordinat, needArr []Coordinat) bool {
	for _,v := range needArr{
	   if need == v{
		   return true
	   }
   }
   return false
}

func inverted(tmp []int) bool {  			//倒序检查
	var i int
	for k := len(tmp)-1; k>=0;k--{  
		if k == 0{  						//最后一个下标说明无法对比
			return false
		}
		if tmp[k]-1 == tmp[k]{  			//说明值是连续数字
			i++
			if i == point{  				//如果达到连续数就返回
				return true
			}
		}else{
			return false
		}
	}
	return false
}

func postive(tmp []int) bool {  			//正序检查
	var i int
	for ck, cv := range tmp {
		if ck == len(tmp)-1{  				//最后一个下标说明无法对比
			return false
		}
		if cv+1 == tmp[ck+1] {  			//说明值是连续数字
			i++
			if i == point{  				//如果达到连续数就返回
				return true
			}
		}else{
			return false
		}
	}
	return false
}

func Slope(inat *Allinat,coor Coordinat) bool {
	var (
		Xmax,Xmin int = coor.x+4,coor.x-4
		Ymax,Ymin int = coor.y+4,coor.y-4
		j,p		int
		lrise,lfall  Coordinat
		//tmp []int
	)
	if Xmin < 0 {
		Xmin = 0
	}
	if Ymin < 0 {
		Ymin = 0
	}
	for i:=1; i<=Xmax-Xmin; i++{
			lrise.x,lrise.y = coor.x,coor.y
			lrise.x = lrise.x+i
			lrise.y = lrise.y+i
		if lrise.x > Xmax || lrise.y > Ymax{
			lrise.x,lrise.y = coor.x,coor.y
			lrise.x = lrise.x-(Xmax-Xmin-i)
			lrise.y = lrise.y-(Xmax-Xmin-i)
			if lrise.x > Xmax || lrise.y < Ymin {
				break
			}
		} 
		if InArray(lrise,inat.key) {
			j++
			if j == (point+1){
				return true
			}
		}
	}
	
	for ii:=1; ii<=Xmax-Xmin; ii++{
			lfall.x,lfall.y = coor.x,coor.y
			lfall.x = lfall.x+ii
			lfall.y = lfall.y-ii
		if lfall.x < Xmin || lfall.y > Ymax {
			lfall.x,lfall.y = coor.x,coor.y
			lfall.x = lfall.x+(Xmax-Xmin-ii)
			lfall.y = lfall.y-(Xmax-Xmin-ii)
			if lfall.x < Xmin || lfall.y < Ymin {
				break
			}
		}
		if InArray(lfall,inat.key) {
			p++
			if p == (point+1){
				return true
			}
		}
	}
	return false
}

func lengthwaysORcrosswise(inat *Allinat,coor Coordinat,S string) bool {
	var (
		max,min int = coor.x+4,coor.x-4
		tmp []int
	)
	if min < 0 {
		min = 0
	}
	for _,c := range inat.key{
		if S == "l" {
			if (max >= c.x && c.y == coor.y) || (min >= c.x  && c.y == coor.y){
				tmp = append(tmp,c.x)
			}
		}else{
			if (max >= c.y && c.x == coor.x) || (min >= c.y  && c.x == coor.x){
				tmp = append(tmp,c.y)
			}
		}
	}
	sort.Ints(tmp)
	if (inverted(tmp) == true) || (postive(tmp) == true) {
		return true
	}
	return false
}

func IsFive(inat *Allinat,coor Coordinat) bool {
	fmt.Println("l:",lengthwaysORcrosswise(inat,coor,"l"))
	fmt.Println("C:",lengthwaysORcrosswise(inat,coor,"C"))
	fmt.Println("Sl",Slope(inat,coor))
	if lengthwaysORcrosswise(inat,coor,"l") == true || lengthwaysORcrosswise(inat,coor,"c") == true || Slope(inat,coor) == true{
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
	fmt.Println(c,"*****",coor)
	if c == false{								//not finish five
		inat.key = append(inat.key,coor)
		fmt.Println("**************没有连成")
		return 
	}
	fmt.Println("**************连成",point+2,"颗")
	return 
	
}

