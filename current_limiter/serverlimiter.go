package main 

import(
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
)
const(
	limiter = 5			
)
var(
	con redis.Conn
	connerr error
)
func init(){
	
	con,connerr = redis.Dial("tcp","127.0.0.1:6379"); if connerr !=nil {
		fmt.Printf("redis连接错误：%v !\n",connerr)
	}
}
func Serlock(ip string){
	var visitime,nowtime int64
	nowtime = time.Now().Add(time.Second * time.Duration(-1)).Unix()
	//con.Do("HSET","limitime",ip,time.Now().Unix())
	val,_ := redis.Values(con.Do("HMGET","limitime",ip))
	redis.Scan(val,&visitime)
	if nowtime-limiter<visitime{
		con.Do("HMGET","limitime",ip)
		con.Send("MULTI")
		con.Do("ZINCRBY",)
	}
	ct := time.Unix(visitime,0)
	fmt.Println("visitime",ct.Format("2006-01-02 15:04:05"))
	bt := time.Unix(nowtime,0)
	fmt.Println("nowtime",bt.Format("2006-01-02 15:04:05"))
}

func main()  {
	Serlock("ads")
}