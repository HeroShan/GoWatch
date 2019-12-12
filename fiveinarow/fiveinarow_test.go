package fiveinarow

import(
	"testing"
)

func TestAddAllinat(t *testing.T){
	 chess := Coordinat{}
	 chess.x = 10
	 chess.y = 15
	 AddCoordinat(chess)
	 chess.x = 10
	 chess.y = 12
	 AddCoordinat(chess)
}