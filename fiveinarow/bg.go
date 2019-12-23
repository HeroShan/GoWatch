package fiveinarow

import(
	 "fmt"
	 "sort"
)
const(
	Vertical = 100
	Level = 100
	matrix = Vertical * Level
	point  =  3
)
type Coordinat struct{
	X		int
	Y		int
	Color	string
}

type Allinat struct{
	Key		[]Coordinat
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
		Xmax,Xmin int = coor.X+4,coor.X-4
		Ymax,Ymin int = coor.Y+4,coor.Y-4
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
			lrise.X,lrise.Y,lrise.Color = coor.X,coor.Y,coor.Color
			lrise.X = lrise.X+i
			lrise.Y = lrise.Y+i
		if lrise.X > Xmax || lrise.Y > Ymax{
			lrise.X,lrise.Y,lrise.Color = coor.X,coor.Y,coor.Color
			lrise.X = lrise.X-(Xmax-Xmin-i)
			lrise.Y = lrise.Y-(Xmax-Xmin-i)
			if lrise.X > Xmax || lrise.Y < Ymin {
				break
			}
		} 
		if InArray(lrise,inat.Key) {
			j++
			if j == (point+1){
				return true
			}
		}
	}
	
	for ii:=1; ii<=Xmax-Xmin; ii++{
			lfall.X,lfall.Y,lfall.Color = coor.X,coor.Y,coor.Color
			lfall.X = lfall.X+ii
			lfall.Y = lfall.Y-ii
		if lfall.X < Xmin || lfall.Y > Ymax {
			lfall.X,lfall.Y,lfall.Color = coor.X,coor.Y,coor.Color
			lfall.X = lfall.X+(Xmax-Xmin-ii)
			lfall.Y = lfall.Y-(Xmax-Xmin-ii)
			if lfall.X < Xmin || lfall.Y < Ymin {
				break
			}
		}
		if InArray(lfall,inat.Key) {
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
		max,min int = coor.X+4,coor.X-4
		tmp []int
	)
	if min < 0 {
		min = 0
	}
	for _,c := range inat.Key{
		if S == "l" {
			if (max >= c.X && c.Y == coor.Y && c.Color == coor.Color) || (min >= c.X  && c.Y == coor.Y && c.Color == coor.Color){
				tmp = append(tmp,c.X)
			}
		}else{
			if (max >= c.Y && c.X == coor.X && c.Color == coor.Color) || (min >= c.Y  && c.X == coor.X && c.Color == coor.Color){
				tmp = append(tmp,c.Y)
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
	if lengthwaysORcrosswise(inat,coor,"l") == true || lengthwaysORcrosswise(inat,coor,"c") == true || Slope(inat,coor) == true{
		return true
	}
	return false
}

func (inat *Allinat)AddCoordinat(coor Coordinat){
	for _,coslice := range inat.Key{
		if coslice.X == coor.X && coslice.Y == coor.Y{
			return 
		}
	}
	c := IsFive(inat,coor)
	fmt.Println(c,"*****",coor)
	if c == false{								//not finish five
		inat.Key = append(inat.Key,coor)
		fmt.Println("**************没有连成")
		return 
	}
	fmt.Println("**************连成",point+2,"颗")
	return 
	
}

