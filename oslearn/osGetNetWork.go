package main

import(
	"fmt"
	"net/http"
)
var Word = []string{"A",
					"B",
					"C",
					"D",
					"E",
					"F",
					"G",
					"H",
					"I",
					"J",
					"K",
					"L",
					"M",
					"N",
					"O",
					"P",
					"Q",
					"R",
					"S",
					"T",
					"U",
					"V",
					"W",
					"X",
					"Y",
					"Z"}

func Getnetwork(url string, chanl chan int){
	url = "http://www."+url+".com"
	_, err := http.Get(url); if err != nil{
		fmt.Println("errrrrrrrrrrr:",url)
	}
	fmt.Println(url)
	<-chanl
	

}

func main(){
	 chanl := make(chan int, 26)
	for four:=0;four<26;four++{
		for three:=0;three<26;three++{
			for two:=0;two<26;two++{
				for one:=0;one<26;one++{
					chanl <- one
					go Getnetwork(Word[four]+Word[three]+Word[two]+Word[one], chanl)
					//fmt.Println(Word[four],Word[three],Word[two],Word[one])
				}
			}	
		}
		
	}
}

