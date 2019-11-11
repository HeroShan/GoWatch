package auth

import (
	_ "net/http"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

// func Islogin(f http.HanderFunc) http.HanderFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.URL.Path)
// 		f(w, r)
// 	}
// }



func Check(username,password string)bool {
	con,connerr := redis.Dial("tcp","47.104.225.152:6379"); if connerr !=nil {
		fmt.Println(connerr)
	}
	infovalue,_ :=redis.Values(con.Do("HMGET","auth",username,password))
	if infovalue[0] != nil && infovalue[1] !=nil{
		return true
	}else{
		return false
	}
}