package main

import(
	"net"
	"fmt"
	"os"
	"strings" 
	"GoWatch/heartbeat"
)


func Getlocalip() (ip string) {
	netInterfaces, err := net.Interfaces()
    if err != nil {
        fmt.Println("net.Interfaces failed, err:", err.Error())
        return ""
	}
 
    for i := 0; i < len(netInterfaces); i++ {
        if (netInterfaces[i].Flags & net.FlagUp) != 0 {
            addrs, _ := netInterfaces[i].Addrs()
 
            for _, address := range addrs {
                if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
                    if ipnet.IP.To4() != nil {
                        return ipnet.IP.String()
                    }
                }
            }
        }
    }
 	return ""
}

func VoteTime(){

}

func HeartBeat(){
	file,err := os.Open("./fsm.config"); if err != nil{
		fmt.Println("file open fail",err)
	}
	data := make([]byte,512)
	var filestr string
	for{
		count, _ := file.Read(data)
		filestr += string(data[:count])
		if count == 0{
			break
		}
	}
	locIp := Getlocalip()
	ip := strings.Split(filestr,"=")
	iplist := strings.Split(ip[1],",")
	for k,v := range iplist{
		v = strings.TrimSpace(v)
		if v == locIp {
			iplist = append(iplist[:k],iplist[k+1:]...)
		}
	}
	go heartbeat.Server()
	heartbeat.Client(iplist)
}


func main(){
	HeartBeat()
	
}
