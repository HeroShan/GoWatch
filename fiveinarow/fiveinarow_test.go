package fiveinarow

import(
	"testing"
	"fmt"
	"math/rand"
	"time"
)

func TestAddAllinat(t *testing.T){
	 c := Allinat{}
	 chess := Coordinat{}
	 for i:=0; i<800000; i++{
		rand.Seed(time.Now().UnixNano())
		chess.x = rand.Intn(50)
		chess.y = rand.Intn(50)
		c.AddCoordinat(chess)
	 }
	//  chess.x = 10
	//  chess.y = 7
	//  c.AddCoordinat(chess)
	// for i:=3; i<=13;i++{
	// 	for j:=6;j>=0;j--{
	// 		chess.x = i
	// 		chess.y = j
	// 		c.AddCoordinat(chess)
	// 	}
	// }
	 for k,v := range c.key{
		 fmt.Println(k,v)
	 }
}

// func TestSlope(t *testing.T){
// 	chess := Coordinat{3,1}
// 	Slope(nil,chess)	
// }