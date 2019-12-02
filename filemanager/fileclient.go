package main

import(
	_"io"
	"fmt"
	"os"
	"bufio"
	"strings"
	"path/filepath"
	"net"
)

func Fileread(str string){
	data := make([]byte,512)
	fstr:=strings.Fields(str)
	for _,Furl := range fstr{
		file ,operr := os.Open(Furl); if operr != nil{
			fmt.Println(Furl,"文件路径错误")
		}else{
			_,filename := filepath.Split(Furl)
			ok,conn := Client([]byte(filename))
			fmt.Println(ok)
			if ok == "ok"{
				for{
					count,_ := file.Read(data); if count == 0{
						break
					}
					n,e := conn.Write(data[:count]);if e != nil{
						fmt.Println(n,e)
					}
					fmt.Println(n)
				}
			}
			
			file.Close()
		}
	}
}

func Cmdput(){
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


func Client(filename []byte)(buf string,udpConn *net.UDPConn){
	udpAddr, _ := net.ResolveUDPAddr("udp", ":8005")

    //连接udpAddr，返回 udpConn
	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
	fmt.Println(filename)
    // 发送数据
	udpConn.Write(filename)
	//读取数据
    buff := make([]byte, 512)
	udpConn.Read(buff)
	buf = string(buff)
	fmt.Println("buff:",buf)
    return buf,udpConn
}

func main(){
	Cmdput()
	
}