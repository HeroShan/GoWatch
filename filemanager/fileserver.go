package main

import (
   "net"
   "fmt"
   "os"
   "io"
)

func RecvFile(fileName string, conn net.Conn)  {
   // 创建新文件
   f, err := os.Create(fileName)
   if err != nil {
      fmt.Println("Create err:", err)
      return
   }
   defer f.Close()

   // 接收客户端发送文件内容，原封不动写入文件
   buf := make([]byte, 4096)
   for {
      n, err := conn.Read(buf)
      if err != nil {
         if err == io.EOF {
            fmt.Println("文件接收完毕")
         } else {
            fmt.Println("Read err:", err)
         }
         return
      }
      f.Write(buf[:n])   // 写入文件，读多少写多少
   }
}

func main()  {
   // 创建监听
   listener, err := net.Listen("tcp", ":1997")
   if err != nil {
      fmt.Println("Listen err:", err)
      return
   }
   defer listener.Close()

   // 阻塞等待客户端连接
   conn, err := listener.Accept()
   if err != nil {
      fmt.Println("Accept err:", err)
      return
   }
   defer conn.Close()

   // 读取客户端发送的文件名
   buf := make([]byte, 1024)
   n, err := conn.Read(buf)
   if err != nil {
      fmt.Println("Read err:", err)
      return
   }
   fileName := string(buf[:n])       // 保存文件名

   // 回复 0k 给发送端
   conn.Write([]byte("ok"))

   // 接收文件内容
   RecvFile(fileName, conn)      // 封装函数接收文件内容， 传fileName 和 conn
}