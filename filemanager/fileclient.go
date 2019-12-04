package main

import (
   "fmt"
   "os"
   "net"
   "io"
   "strings"
   "sync"
)

func SendFile(path string, conn net.Conn)  {
   // 以只读方式打开文件
   f, err := os.Open(path)
   if err != nil {
      fmt.Println("os.Open err:", err)
      return
   }
   defer func(){
      f.Close()
   }()                   // 发送结束关闭文件。

   // 循环读取文件，原封不动的写给服务器
   buf := make([]byte, 4096)
   for {
      n, err := f.Read(buf)        // 读取文件内容到切片缓冲中
      if err != nil {
         if err == io.EOF {
            fmt.Println(path,"文件发送完毕")
            return 
         } else {
            fmt.Println(path,":发送失败")
            return 
         }
         return
      }
      conn.Write(buf[:n])  // 原封不动写给服务器
   }

}


func FielCh(path string,conn net.Conn){
   // 获取文件名   fileInfo.Name()
   fileInfo, err := os.Stat(path)
   if err != nil {
      fmt.Println("os.Stat err:", err)
      return
   }

   // 主动连接服务器
   

   // 给接收端，先发送文件名
   _, err = conn.Write([]byte(fileInfo.Name()))
   if err != nil {
      fmt.Println("conn.Write err:", err)
      return
   }

   // 读取接收端回发确认数据 —— ok
   buf := make([]byte, 1024)
   n, err := conn.Read(buf)
   if err != nil {
      fmt.Println("conn.Read err:", err)
      return
   }

   // 判断如果是ok，则发送文件内容
   if "ok" == string(buf[:n]) {
      SendFile(path, conn)   // 封装函数读文件，发送给服务器，需要path、conn
   }
}

func main()  {
   // 提示输入文件名
      var lock  sync.Mutex

      conn, err := net.Dial("tcp", "47.104.225.152:1997")
      if err != nil {
         fmt.Println("net.Dial err:", err)
         return
      }
      fmt.Println("请输入需要传输的文件(多文件用“，”隔开)：")
      var path string
      fmt.Scan(&path)
      file := strings.Split(path,",");
      for _,fpath := range file{
         lock.Lock()
         FielCh(fpath,conn)
         lock.Unlock()
      }

}