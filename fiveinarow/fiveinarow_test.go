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
	 chess.x = 10
	 chess.y = 12
	 c.AddCoordinat(chess)
	 chess.x = 18
	 chess.y = 12
	 c.AddCoordinat(chess)
	 chess.x = 18
	 chess.y = 12
	 c.AddCoordinat(chess)
	 fmt.Println(c)
}