package main

import(
	_"io"
	"fmt"
	"os"
	"bufio"
	"strings"
)

func Fileread(str string){
	fstr:=strings.Fields(str)
	for k,v := range fstr{
		fmt.Println(k,v)
	}
}

func main(){
	var command string
	for{
		fmt.Printf("请输入命令：")
		fmt.Scanf("%v\n",&command)
		switch command {
		case "upload":
			fmt.Printf("请输入文件路径，用空格隔开：")
			inputReader := bufio.NewReader(os.Stdin)
			str,_ := inputReader.ReadString('\n')
			Fileread(str)
		
		case "download":
			fmt.Println(command)	
		
		case "exit":
			os.Exit(0)	
		}
	}
	
}