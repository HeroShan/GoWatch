package main
import(
	"fmt"
	"math/rand"
	"time"
	"GoWatch/fiveinarow"
)
func main(){
	c := fiveinarow.Allinat{}
	 chess := fiveinarow.Coordinat{}
	 for i:=0; i<10000; i++{
		if i % 2 == 0{
			chess.Color = "white"
		}else{
			chess.Color = "black"
		}
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