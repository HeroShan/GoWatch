package mapapi

import(
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
	
)
var isset int64 = 0
func Connect(Ip string, Path string, Point JsPoint) {
	con,connerr := redis.Dial("tcp","127.0.0.1:6379"); if connerr !=nil {
		fmt.Println(connerr)
	}
	timeUnix:=time.Now().Unix()   //已知的时间戳
	now :=time.Unix(timeUnix,0).Format("2006-01-02 15:04:05")
	issetIp,_ := con.Do("SISMEMBER","ip",Ip)
	if issetIp == isset {
		con.Do("SADD","ip",Ip)
		con.Do("SADD","ipArea",Ip+"~"+Path+"~"+now+"~"+Point.X+"~"+Point.Y)
	}
	
}