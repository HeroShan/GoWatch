package fiveinarow

import(
	"testing"
	_"fmt"
)

func TestAddAllinat(t *testing.T){
	 c := Allinat{}
	 chess := Coordinat{}
	 chess.x = 10
	 chess.y = 15
	 c.AddCoordinat(chess)
	 chess.x = 0
	 for i:=0; i<100; i ++{
		chess.y = i
		c.AddCoordinat(chess)
	 }
	 chess.y = 11
	 for i:=0; i<100; i ++{
		chess.x = i
		c.AddCoordinat(chess)
	 }
}