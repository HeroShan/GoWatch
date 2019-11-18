package heartbeat
import(
	"github.com/dgrijalva/jwt-go"
	"encoding/base64" 

)
var Key string = "nR5cCI6I"
func Auth()string{
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString([]byte(Key))
	return base64.StdEncoding.EncodeToString([]byte(tokenString))
	// dd,_:=base64.StdEncoding.DecodeString(cc)
	// fmt.Println(tokenString)
	// fmt.Println(cc)
	// fmt.Println(string(dd))
}



