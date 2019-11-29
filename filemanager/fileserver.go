package main

import(
	"net"
	"fmt"
    "io"
    "os"
)

func SaveFile(fileName string,udpConn *net.UDPConn) {
    file,ferr := os.Create("./File/"+fileName); if ferr != nil{
        fmt.Println("file----",ferr)
    }
    fmt.Println(udpConn)
    // 读取数据
    buf := make([]byte, 512)
    len, err := udpConn.Read(buf)
    file.Write(buf[:len])
    if err != nil{
        if err == io.EOF{
            fmt.Println("UDP read err:",err)  
            file.Close()
        }else{
            fmt.Println("UDP read err:",err)  
        }
    }
    
    
}

func main(){
	udpAddr,_ := net.ResolveUDPAddr("udp","127.0.0.1:1997")

	udpConn, err := net.ListenUDP("udp", udpAddr)
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
    _, err = udpConn.WriteToUDP([]byte("ok"), udpAddr)
    fileName := string(buf[:n]) 
    SaveFile(fileName,udpConn)


}

