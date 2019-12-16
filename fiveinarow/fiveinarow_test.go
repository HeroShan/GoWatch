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
	//  chess.x = 10
	//  chess.y = 15
	//  c.AddCoordinat(chess)
	 
	 for i:=0; i<20000; i++{
		rand.Seed(time.Now().UnixNano())
		chess.x = rand.Intn(9)
		chess.y = rand.Intn(9)
		c.AddCoordinat(chess)
		
	 }
	 for k,v := range c.key{
		 fmt.Println(k,v)
	 }
}