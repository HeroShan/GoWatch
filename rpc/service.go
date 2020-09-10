package rpc
import(
	"net/rpc"
	"net"
	"net/http"
	"log"
)
type ZooReq struct{}

type ZooRes struct{
	Say string
}

type Zoo struct{
	Monkey	string
	Cow		string
}

func(ZooReq *ZooReq) MonkeyEcho(Z *Zoo,Zr *ZooRes) error {
	if Z.Monkey != ""{
		Zr.Say =  Z.Monkey+":~hohoho!"
	}else{
		Zr.Say = "you come out!"
	}
	log.Printf("Rpc responseis :%#v!",Zr)
	return nil
}
func(ZooReq *ZooReq) CowEcho(Z *Zoo,Zr *ZooRes) error {
	if Z.Cow != ""{
		Zr.Say =  Z.Cow+":~momomo!"
	}else{
		Zr.Say = "you come back!"
	}
	log.Printf("Rpc responseis :%#v!",Zr)
	return nil
}

func Service(){
	rpc.Register(new(ZooReq))
	rpc.HandleHTTP()
	l, _ := net.Listen("tcp", ":1997")
	http.Serve(l, nil)
}

