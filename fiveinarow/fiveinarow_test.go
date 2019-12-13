package fiveinarow

import(
	"testing"
	"fmt"
)

func TestAddAllinat(t *testing.T){
	 c := Allinat{}
	 chess := Coordinat{}
	 chess.x = 10
	 chess.y = 15
	 c.AddCoordinat(chess)
	 chess.x = 0
	 for i:=0; i<10; i ++{
		chess.y = i
		c.AddCoordinat(chess)
	 }
	 fmt.Println(c)
}