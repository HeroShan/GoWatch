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
	 for i:=0; i<8000; i++{
		rand.Seed(time.Now().UnixNano())
		chess.X = rand.Intn(10)
		chess.Y = rand.Intn(10)
		c.AddCoordinat(chess)
	 }
	//  chess.X = 10
	//  chess.Y = 7
	//  c.AddCoordinat(chess)
	// for i:=3; i<=13;i++{
	// 	for j:=6;j>=0;j--{
	// 		chess.X = i
	// 		chess.Y = j
	// 		c.AddCoordinat(chess)
	// 	}
	// }
	 for k,v := range c.Key{
		 fmt.Println(k,v)
	 }
}

// func TestSlope(t *testing.T){
// 	chess := Coordinat{3,1}
// 	Slope(nil,chess)	
// }