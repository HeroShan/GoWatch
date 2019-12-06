package main

import(
	"os"
	"fmt"
)

type PathPrama struct{
	Path 	[]string
}



func main(){
	var P 	PathPrama
	P.Path	= os.Args
	fmt.Println(P)
}