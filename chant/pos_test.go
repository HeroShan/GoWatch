package chant

import(
	"testing"
	"fmt"
)

func TestSearchNum(t *testing.T){
	var p PSOParam
	p.Large = 20
	p.ParticleNum = 1
	p.Matrix = 3
	p.Step   = p.Large / p.ParticleNum
	var i int
	
		particle := p.CreateParticle()
		fmt.Println("particle :",particle)
		mapp := p.CreateMap()
		fmt.Println("map :",mapp)
		for{	
			i++
			// fmt.Printf("请输入查找数字:")
			// fmt.Scanf("%d\n",&targetNum)
			c,cc := p.SearchNum(mapp,&particle,i)
			fmt.Printf("Num is %d ,Search is : %d,index is %d\n",i,c,cc)
			if i== 20{
				break
			}
		}
}