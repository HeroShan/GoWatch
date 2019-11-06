package main

import(
	"os"
	"fmt"
	"bufio"
)

func OpenFile(File string) {
	f,err := os.Open(File);	if err !=nil {
		fmt.Println(err)
	}
	data := make([]byte, 1024)
	fcrd := bufio.NewReader(f)
	for{
		_,err := fcrd.Read(data);	if err !=nil {
			fmt.Println("errrrrrr",err)
			break
		}
		fmt.Println(string(data))
	}
	//fmt.Println(f)
	defer f.Close()
}

var File string

func main(){
	File = "../css/img/love.jpg"
	OpenFile(File)

}
