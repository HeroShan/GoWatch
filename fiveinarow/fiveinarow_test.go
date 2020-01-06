package fiveinarow

import(
	"testing"
	"fmt"
	"math/rand"
	"time"
)

func Benchmark_AddAllinat(b *testing.B){
	for k:=0;k<b.N;k++{
		c := Allinat{}
	    chess := Coordinat{}
		for i:=0; i<900; i++{
			if i % 2 == 0{
				chess.Color = "white"
			}else{
				chess.Color = "black"
			}
			rand.Seed(time.Now().UnixNano())
			chess.X = rand.Intn(5)
			chess.Y = rand.Intn(5)
			xx,x := c.AddCoordinat(chess)
			if xx == true{
				fmt.Printf("%#v\n\n",x)
				break
			}
			
		}
	}
}


func TestAddAllinat(t *testing.T){
	c := Allinat{}
	chess := Coordinat{}
	for i:=0; i<500000; i++{
	   if i % 2 == 0{
		   chess.Color = "white"
	   }else{
		   chess.Color = "black"
	   }
	   rand.Seed(time.Now().UnixNano())
		   chess.X = rand.Intn(50)
		   chess.Y = rand.Intn(50)
		   xx,x := c.AddCoordinat(chess)
		   if xx == true{
			   fmt.Println(x,chess)
		   }
	   
	}

}



// func TestAddAllinat(t *testing.T){
// 	 c := Allinat{}
// 	 chess := Coordinat{}
// 	 cch := make(chan int,4)
// 	 for i:=0; i<50000000; i++{
// 		if i % 2 == 0{
// 			chess.Color = "white"
// 		}else{
// 			chess.Color = "black"
// 		}
// 		cch <- i
// 		go func(ccs chan int){
// 			rand.Seed(time.Now().UnixNano())
// 			chess.X = rand.Intn(130)
// 			chess.Y = rand.Intn(130)
// 			xx,x := c.AddCoordinat(chess)
// 			if xx == true{
// 				fmt.Printf("%#v===========%#v\n\n",x,chess)
// 			}
// 			<-ccs
// 		}(cch)
		
// 	 }
// 	//  chess.X = 10
// 	//  chess.Y = 7
// 	//  c.AddCoordinat(chess)
// 	// for i:=3; i<=13;i++{
// 	// 	for j:=6;j>=0;j--{
// 	// 		chess.X = i
// 	// 		chess.Y = j
// 	// 		c.AddCoordinat(chess)
// 	// 	}
// 	// }
// 	//  for k,v := range c.Key{
// 	// 	 fmt.Println(k,v)
// 	//  }
// }

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