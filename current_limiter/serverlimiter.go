package current_limiter

import(
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
)
const(
	timeMax = 5			//time is second
	limiter = 5			//limit count
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
func Serlock(ip string) bool {		//根据IP进行限流
	var(
		visitime,nowtime int64
	)
	nowtime = time.Now().Unix()
	val,_ := redis.Values(con.Do("HMGET","limitime",ip))
	redis.Scan(val,&visitime)
	if visitime == 0{
		con.Do("HMSET","limitime",ip,nowtime)
		return true
	}
	if nowtime-timeMax<visitime{
		redint,_ := redis.Int(con.Do("ZSCORE","limitimescore",ip))
		if redint >= limiter{
			return false
		}
		con.Send("MULTI")								//开启事物
		con.Do("ZINCRBY","limitimescore",1,ip)			//频率限制+1最大数为5
		con.Do("HMSET","limitime",ip,nowtime)			//设置当前的时间
		_,err := con.Do("EXEC")							//提交事物
		if err != nil{
			fmt.Println(err)
			return false
		}
	}
	con.Do("HMSET","limitime",ip,nowtime)
	return true
}


