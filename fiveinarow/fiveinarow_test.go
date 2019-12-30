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
	 for i:=0; i<900000; i++{
		if i % 2 == 0{
			chess.Color = "white"
		}else{
			chess.Color = "black"
		}
		rand.Seed(time.Now().UnixNano())
		chess.X = rand.Intn(100)
		chess.Y = rand.Intn(100)
		xx,x := c.AddCoordinat(chess)
		if xx == true{
			fmt.Printf("%#v\n\n",x)
			break
		}
		
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
	//  for k,v := range c.Key{
	// 	 fmt.Println(k,v)
	//  }
}

// func TestSlope(t *testing.T){
// 	chess := Coordinat{3,1}
// 	Slope(nil,chess)	
// }

// func TestCenter(t *testing.T){
// 	a := new(Allinat)
// 	chess := Coordinat{3,1,""}
// 	a.AddCoordinat(chess)
// 	a.Center()
// 	fmt.Println(a.Key)
// }