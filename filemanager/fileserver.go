package main

import(
	"net"
	"fmt"
    "io"
    "os"
)

func SaveFile(fileName string,udpConn net.Conn) {
    file,_ := os.Create("./File/"+fileName)
    
    // 读取数据
    buf := make([]byte, 512)
    for {
        n, err := udpConn.Read(buf)
        if err != nil {
           if err == io.EOF {
              fmt.Println("文件接收完毕")
           } else {
              fmt.Println("Read err:", err)
           }
           return
        }
        fmt.Println("fiel-----------------",string(buf[:n]))
        file.Write(buf[:n])   // 写入文件，读多少写多少
    }
    
    
}

func main(){
	listener, err := net.Listen("tcp", "127.0.0.1:8005")
    if err != nil {
        fmt.Println("Listen err:", err)
        return
    }
    defer listener.Close()

    // 阻塞等待客户端连接
    udpConn, err := listener.Accept()
	if err != nil {
        fmt.Println(err)
    }
    //defer udpConn.Close()
    
    buf := make([]byte, 1024)
    n, err := udpConn.Read(buf)
    if err != nil {
        fmt.Println("Read err:", err)
        return
    }
    _, err = udpConn.Write([]byte("ok"))
    if err != nil{
        fmt.Println("write-------",err)
    }
    fileName := string(buf[:n]) 
    SaveFile(fileName,udpConn)


}

