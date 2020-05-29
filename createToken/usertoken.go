package createToken

import(
	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
	"time"
	"fmt"
	"strconv"
)


//	Get Token for cookie     and add redis to hash Map    expire time is 7 Days
func GetToken(day int)string {
	createTime := time.Now().AddDate(0,0,day).UnixNano()		//秒为单位  int64 过期时间7day
	var LoSecretKey string = strconv.FormatInt(createTime,10)
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString([]byte(LoSecretKey))
	con,connerr := redis.Dial("tcp","127.0.0.1:6379"); if connerr !=nil {
		fmt.Println(connerr)
	}
	con.Do("HMSET","loginToken",tokenString,LoSecretKey)
	con.Close()
	return tokenString
}

//	Get Cookie compare Redis loginToken-Value 
//	return false
func IsLogin(Wisheart string) int64 {
	createTime := time.Now().UnixNano()
	con,connerr := redis.Dial("tcp","127.0.0.1:6379"); if connerr !=nil {
		fmt.Println(connerr)
	}
	r,_ := redis.Int64(con.Do("hget","loginToken",Wisheart))
	expiretime := r-createTime
	con.Close()
	return expiretime
	
}

//	Function is  Delete expire seven days Redis login Token
func DelExpireToken(){
	con,connerr := redis.Dial("tcp","127.0.0.1:6379"); if connerr !=nil {
		fmt.Println(connerr)
	}
	createTime := time.Now().UnixNano()
	r,_:=redis.StringMap(con.Do("hgetall","loginToken"))
	for k,v := range r{
		vint64,_:=strconv.ParseInt(v, 10, 64)
		expireToken := vint64-createTime
		if expireToken<0{
			con.Do("hdel","loginToken",k,v)
		}
	}
}