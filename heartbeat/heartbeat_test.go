package heartbeat

import(
	"testing"
)


func BenchmarkClient(b *testing.B){
	var iplist = []string{"183.159.99.163","47.105.217.180","47.104.225.152"}
	for i:=0;i<b.N;i++{
	Client(iplist,"125.119.249.172")
	}
}
