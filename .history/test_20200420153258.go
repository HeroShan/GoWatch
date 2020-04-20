package main

import(
	"fmt"
	"runtime"
)

func main(){
	c :=runtime.NumCPU()
	fmt.Println(c)
}