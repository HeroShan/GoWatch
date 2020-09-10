package rpc
import (
    "fmt"
    "log"
    "net/rpc"
)

func Client(){
	conn, err := rpc.DialHTTP("tcp", "47.104.225.152:1997")
    if err != nil {
        log.Fatalln("dailing error: ", err)
	}
	var z Zoo
	var zs ZooRes
	z.Monkey = "heihei"
	go func(){
		z.Cow = "baibai"
		err = conn.Call("ZooReq.CowEcho",&z,&zs)
		if err != nil {
			log.Fatalln("arith error: ", err)
		}
		fmt.Printf("rpc:%#v\n",zs)
	}()
	err = conn.Call("ZooReq.MonkeyEcho",&z,&zs)
	if err != nil {
        log.Fatalln("arith error: ", err)
	}
	
}

