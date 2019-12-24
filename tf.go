package main
import(
	"fmt"
	"math/rand"
	"time"
	"GoWatch/fiveinarow"
)
func main(){
	c := Allinat{}
	chess := Coordinat{}
	for i:=0; i<88888; i++{
	   if i % 2 == 0{
		   chess.Color = "white"
	   }else{
		   chess.Color = "black"
	   }
	   rand.Seed(time.Now().UnixNano())
	   chess.X = rand.Intn(8)
	   chess.Y = rand.Intn(9)
	   x,xarr := c.AddCoordinat(chess)
	   if x == true {
		   fmt.Println("true",xarr)
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
	for k,v := range c.Key{
		fmt.Println(k,v)
	}
}