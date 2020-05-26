package ws

import(
	"testing"
	"fmt"
)

type TestPro struct {
    msgContent   string
}

// 实现发送者
func (p *TestPro) MsgContent() string {
    return p.msgContent
}

// 实现接收者
func (t *TestPro) Consumer(dataByte []byte) error {
    fmt.Println(string(dataByte))
    return nil
}

func TestSend(t *testing.T){
	c := Conn()
	sendrespons := Send(c,"wsMessage",[]byte("this is a test message!"))
	c1 := Conn()
	Send(c1,"wsMessage",[]byte("this is a test message!"))
	fmt.Printf("%v\n",sendrespons)
	c2 := Conn()
	Receive(c2,"wsMessage")
	

}