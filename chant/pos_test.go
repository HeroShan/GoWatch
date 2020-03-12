package chant

import(
	"testing"
	"fmt"
	"time"
	"math/rand"
)

func TestSearchNum(t *testing.T){
	p  := &PSOParam{}
	p.Large = 15
	p.ParticleNum = 2
	p.Matrix = 3
	p.Step   = p.Large / p.ParticleNum
	fmt.Printf("%#v \n",p)
	var i int
	
		particle := p.CreateParticle()
		fmt.Println("particle :",particle)
		mapp := p.CreateMap()
		fmt.Printf("map : %#v\n",&mapp)
		rand.Seed(time.Now().Unix())
		for{	
			i++
			// fmt.Printf("请输入查找数字:")
			// fmt.Scanf("%d\n",&targetNum)
			
			q:=rand.Intn(10)
			c,cc := p.SearchNum(mapp[:],&particle,q)
			fmt.Printf("Num is %d ,Search is : %d,index is %d\n",q,c,cc)
			if i== 10{
				break
			}
		}
}

