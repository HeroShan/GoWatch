package main

import(
	"fmt"
	"net/http"
	"time"
	_ "sync"
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
var t3 int64
var t1,t2 int64 =  1,1

func Getnetwork(url string, chanl chan int){
	
	url = "http://www."+url+".com"
	_, err := http.Get(url); if err != nil{
		fmt.Println("errrrrrrrrrrr:",url)
	}
	fmt.Println(url)
	<-chanl
	
	
}

func G2etnetwork(url string, chanl2 chan int){
	
	url = "http://www."+url+".com"
	_, err := http.Get(url); if err != nil{
		fmt.Println("errrrrrrrrrrr:",url)
	}
	fmt.Println(url)
	<-chanl2
	
}
func ex1(){
	
	chanl := make(chan int, 300)
	for four:=0;four<26;four++{
		for three:=0;three<13;three++{
			for two:=0;two<26;two++{
				for one:=0;one<26;one++{
					chanl <- one
					go Getnetwork(Word[four]+Word[three]+Word[two]+Word[one], chanl)
					fmt.Println(Word[four],Word[three],Word[two],Word[one])
				}
			}	
		}
		if four ==25{
			t2 =  time.Now().UnixNano()
		}
	}
	defer close(chanl)
}

func ex2(){
	
	chanl2 := make(chan int, 300)
	for four:=0;four<26;four++{
		for three:=13;three<26;three++{
			for two:=0;two<26;two++{
				for one:=0;one<26;one++{
					chanl2 <- one
					go G2etnetwork(Word[four]+Word[three]+Word[two]+Word[one], chanl2)
					fmt.Println(Word[four],Word[three],Word[two],Word[one])
				}
			}	
		}
		if four ==25{
			t1 = time.Now().UnixNano()
		}
	}
	defer close(chanl2)
}

func init(){
	t3 = time.Now().UnixNano()
}
func main(){
	defer func(t3,t1,t2 int64){
		fmt.Println(t3,t1,t2)
	}(t3,t1,t2)
	go ex1()
	ex2()
	for{
		//fmt.Println(&t1,&t2)
		if t1 !=int64(1) && t2 !=int64(1){
			fmt.Println(t1,t2)
			break
		}	
	}
	
}

