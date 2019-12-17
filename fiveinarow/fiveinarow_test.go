package fiveinarow

import(
	"testing"
	_"fmt"
	"math/rand"
	"time"
)

func TestAddAllinat(t *testing.T){
	 c := Allinat{}
	 chess := Coordinat{}
	 chess.x = 10
	 chess.y = 15
	 c.AddCoordinat(chess)
	 
	 for i:=0; i<1150000; i++{
		rand.Seed(time.Now().UnixNano())
		chess.x = rand.Intn(50)
		chess.y = rand.Intn(50)
		c.AddCoordinat(chess)
	 }
	// for i:=0; i<=13;i=i+2{
	// 	for j:=6;j>=0;j--{
	// 		chess.x = i
	// 		chess.y = j
	// 		c.AddCoordinat(chess)
	// 	}
	// }
	//  for k,v := range c.key{
	// 	 fmt.Println(k,v)
	//  }
}

func TestSlope(t *testing.T){
	chess := Coordinat{3,1}
	Slope(nil,chess)	
}