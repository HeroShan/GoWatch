package heartbeat

import(
	"testing"
)
func TestClient(t *testing.T){
	msg := "127.0.0.1"
	Sendmsg(msg)
}