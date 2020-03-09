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
var ParticleSwarm []string{}
type Particle interface{}		//粒子

	

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

func (Pso PSOParam)CreateParticle() ParticleSwarm {		//创建粒子群
	 var(
		blob 	  int
		i 		  int
	 )
	 blob = Pso.Large / Pso.ParticleNum
	 for i = 1; i < blob; i++{
		ParticleSwarm[i]["part"+strconv.Itoa(i)] = Particle
	 }
	return ParticleSwarm


}

func (Pso PSOParam)GetRound(){


}

func main(){
	var p PSOParam
	p.Large = 2000
	p.ParticleNum = 3
	p.Matrix = 3

		//map := p.CreateMap()
		particle := CreateParticle()
		fmt.Println(particle)
}