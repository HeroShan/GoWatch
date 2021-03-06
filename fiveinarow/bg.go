package fiveinarow

import(
	"sort"
	"math"
)
const(
	Vertical = 100
	Level = 100
	matrix = Vertical * Level
	point  =  4
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
	   if v == need {
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
			i = 0
		}
	}
	return false
}

func postive(tmp []int) (bool,int) {  			//正序检查
	var i int
	for ck, cv := range tmp {
		if ck == len(tmp)-1{  				//最后一个下标说明无法对比
			return false,-1
		}
		if cv+1 == tmp[ck+1] {  			//说明值是连续数字
			i++
			if i == point{  				//如果达到连续数就返回
				return true,ck+1
			}
		}else{
			i = 0
		}
	}
	return false,-1
}

func Slope(inat *Allinat,coor Coordinat) (bool,[]Coordinat) {
	var (
		Xmax,Xmin int = coor.X+4,coor.X-4
		Ymax,Ymin int = coor.Y+4,coor.Y-4
		j,p		int
		lrise,lfall  Coordinat
		lrtmp,lftmp []Coordinat
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
			lrise.X = Xmax-i
			lrise.Y = Ymax-i
			if lrise.X > Xmax || lrise.Y < Ymin {
				continue
			}
		} 
		if lrise.X == Xmax || lrise.Y==Ymax {
			if j > 0 && j != point {
				j = 0
				lrtmp = []Coordinat{}
			}
		}
		if InArray(lrise,inat.Key) {
			lrtmp = append(lrtmp,lrise)
			j++
			if j == point{
				lrtmp = append(lrtmp,coor)
				return true,lrtmp
			}
		}else{
			if i>=1 && i<= (Xmax-Xmin)/2 {
				if Xmax < 5{
					continue
				}
				if coor.X+i>=Xmax && j<point{
					j = 0
					lrtmp = []Coordinat{}
					continue
				}
					i = int(math.Ceil(float64((Xmax-Xmin)/2)))
					lrise.X = coor.X-1
					lrise.Y = coor.Y-1
					if !InArray(lrise,inat.Key){
						j = 0
						lrtmp = []Coordinat{}
					}
					continue
			}
			j = 0
			lrtmp = []Coordinat{}
		}
	}
	
	for ii:=1; ii<=Xmax-Xmin; ii++{
			lfall.X,lfall.Y,lfall.Color = coor.X,coor.Y,coor.Color
			lfall.X = lfall.X+ii
			lfall.Y = lfall.Y-ii
		if lfall.X > Xmax || lfall.Y < Ymin {
			lfall.X,lfall.Y,lfall.Color = coor.X,coor.Y,coor.Color
			lfall.X = Xmax-ii
			lfall.Y = Ymax-Ymin+ii
			if lfall.X > Xmin || lfall.Y > Ymax {
				continue
			}
		}
		if InArray(lfall,inat.Key) {
			lftmp = append(lftmp,lfall)
			p++
			if p == point{
				lftmp = append(lftmp,coor)
				return true,lftmp
			}
		}else{
			if ii>=1 && ii<= (Xmax-Xmin)/2 {
				if Xmax < 5{
					continue
				}
					ii = int(math.Ceil(float64((Xmax-Xmin)/2)))
					lfall.X = coor.X+1
					lfall.Y = coor.Y-1
					if !InArray(lfall,inat.Key){
						p = 0
						lftmp = []Coordinat{}
						continue
					}
				if (coor.X+ii)>=Xmax && p<point{
					p = 0
					lftmp = []Coordinat{}
				}
				continue
			}
			p = 0
			lftmp = []Coordinat{}
		}
	}
	return false,[]Coordinat{}
}

func lengthwaysORcrosswise(inat *Allinat,coor Coordinat,S string) (bool,[]Coordinat) {
	var (
		max,min int = coor.X+4,coor.X-4
		tmp 	[]int
		lcinat  []Coordinat
	)
	if min < 0 {
		min = 0
	}
	for _,c := range inat.Key{
		if S == "c" {
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
	pbool ,pint := postive(tmp)
	if pbool == true {
		ci := pint
		for pi := 0; pi <= point; pi++{
			if S == "c"{
				lcinat = append(lcinat,Coordinat{tmp[ci],coor.Y,coor.Color})
			}else{
				lcinat = append(lcinat,Coordinat{coor.X,tmp[ci],coor.Color})
			}
			ci--
		}
		return true,lcinat
	}else{
		return false,lcinat
	}
	// ibool ,iint := inverted(tmp)
	// if (inverted(tmp) == true) || (postive(tmp) == true) {
	// 	return true
	// }
	// return false
}

func IsFive(inat *Allinat,coor Coordinat) (bool,[]Coordinat) {
	// if lengthwaysORcrosswise(inat,coor,"l") == true || lengthwaysORcrosswise(inat,coor,"c") == true || Slope(inat,coor) == true{
	// 	return true
	// }
	// return false
	sis,sarr := Slope(inat,coor);if sis == true{
		return true,sarr
	}
	lis,larr := lengthwaysORcrosswise(inat,coor,"l"); if lis == true{
		return true,larr
	}
	cis,carr := lengthwaysORcrosswise(inat,coor,"c"); if cis == true{
		return true,carr
	}else{
		return false, []Coordinat{}
	}
	
	
	
	
	
}

func (inat *Allinat)AddCoordinat(coor Coordinat) (bool,[]Coordinat) {
	for _,coslice := range inat.Key{
		if coslice.X == coor.X && coslice.Y == coor.Y{
			return false, []Coordinat{}
		}
	}
	c,allfive := IsFive(inat,coor)
	if c == false{								//not finish five
		inat.Key = append(inat.Key,coor)
		return false, []Coordinat{}
	}
	return true,allfive
	
}

