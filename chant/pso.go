package main

import(
	//"time"
	//"math/rand"
	"fmt"
	"strconv"
)

type PSOParam struct{
	Large			int			//地图大小
	ParticleNum		int			//粒子数量
	Matrix			int			//探索矩阵大小
}

type Particle struct{
	 name		string
	 index		int
}		//粒子

	

func (Pso PSOParam)CreateMap() []int{		//创建地图
	var(
		Map  []int
		i	int
	)
	for i = 0;i < Pso.Large; i++{
		Map = append(Map,i)
	}
	return Map 

}

func (Pso PSOParam)CreateParticle()(int , []Particle){		//创建粒子群
	var(
		i 		  			int
		ParticleSwarm  	=   make([]Particle,Pso.ParticleNum)
		step				int
	 )
	 step = Pso.Large / Pso.ParticleNum
	 for i = 1 ; i <= Pso.ParticleNum; i++{
		ParticleSwarm[i-1].name 	= "part"+strconv.Itoa(i)
		if i == 1{
			ParticleSwarm[i-1].index	= step 
		}
		ParticleSwarm[i-1].index	= step * (i-1)
	 }
	  
	return (Pso.Large % Pso.ParticleNum) , ParticleSwarm
}

func (Pso PSOParam)GetRound(){


}

func main(){
	var p PSOParam
	p.Large = 2000
	p.ParticleNum = 15
	p.Matrix = 3

		//map := p.CreateMap()
		mod,particle := p.CreateParticle()
		fmt.Println("mod :",mod)
		for k,v := range particle{
			fmt.Printf("%d , %#v \n",k,v)
		}
}