package main

import(
	_"io"
	"fmt"
	"os"
	"bufio"
	"strings"
	"path/filepath"
)

func Fileread(str string){
	data := make([]byte,512)
	fstr:=strings.Fields(str)
	for _,Furl := range fstr{
		file ,operr := os.Open(Furl); if operr != nil{
			fmt.Println(Furl,"文件路径错误")
		}
		_,filename := filepath.Split(Furl)
		newfile,crerr := os.Create("./"+filename); if crerr != nil{
			fmt.Println("文件名重复")
		}
		for{
			count,_ := file.Read(data); if count == 0{
				break
			}
			newfile.Write(data[:count])
			
		}
		file.Close()
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