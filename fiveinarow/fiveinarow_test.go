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
	 chess.x = 10
	 chess.y = 15
	 c.AddCoordinat(chess)
	 
	 for i:=0; i<115000; i++{
		rand.Seed(time.Now().UnixNano())
		chess.x = rand.Intn(20)
		chess.y = rand.Intn(20)
		c.AddCoordinat(chess)
	 }
	// for i:=0; i<=13;i=i+2{
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