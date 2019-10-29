package main
import(
	"github.com/juju/errors"
	"GoWatch/getuuid"
	e "errors"
	"fmt"
	"GoWatch/mapapi"
)

func te() string{
	var i int
	var cr string
	i ++
	if i>3{
	ter :=	e.New("number out to range")
	fmt.Println(ter)
	cr = errors.ErrorStack(ter)
	}
	return cr
	
	
}
func uuid(){
    ui :=	getuuid.GetV4()
	fmt.Println(ui)
}

var snum int
var ip string
var Point mapapi.JsPoint
func main(){
	ip = "125.121.244.70"
	Point = mapapi.Getpoint(ip)
	fmt.Println(Point)
}