package createToken

import(
	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
	"time"
	"fmt"
	"strconv"
)


//	GET Token for cookie     and add redis to hash Map    expire time is 7 Days
func GetToken()string {
	createTime := time.Now().AddDate(0,0,7).UnixNano()		//秒为单位  int64 过期时间7day
	var LoSecretKey string = strconv.FormatInt(createTime,10)
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString([]byte(LoSecretKey))
	con,connerr := redis.Dial("tcp","47.104.225.152:6379"); if connerr !=nil {
		fmt.Println(connerr)
	}
	con.Do("HMSET","loginToken",tokenString,LoSecretKey)
	return tokenString
}

func IsLogin(Wisheart string) {
	con,connerr := redis.Dial("tcp","47.104.225.152:6379"); if connerr !=nil {
		fmt.Println(connerr)
	}
	r,e := con.Do("HGET","loginToken",Wisheart)
	fmt.Println(r,e)
	fmt.Printf("%T,%T",r,e)

}