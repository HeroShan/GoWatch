package main
import(
	"github.com/juju/errors"
	"GoWatch/getuuid"
	e "errors"
	"fmt"
	"GoWatch/mapapi"
	"GoWatch/seckill"
	"net/http"
	"strings"
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
var areainfo string
var path string
func getip(w http.ResponseWriter, r *http.Request){
	ipstring := r.RemoteAddr
	ip1 := strings.FieldsFunc(ipstring,Splitstr)
	ip = ip1[0]
	//ip = "180.101.49.11"
	//ip = "183.131.107.149"
	Point,areainfo = mapapi.Getpoint(ip)
	if Point.X == "" {
		//fmt.Println("IP format incorrect")
		w.Write([]byte("IP format incorrect"))
	}
	//Point.X = "32.05723550"
	//Point.Y = "120.21287663"
	area := mapapi.Getarea(Point)
	if area != "" {
		path = area
	}else{
		path = areainfo
	}
	w.Write([]byte(path))
	mapapi.Connect(ip,path,Point)
}

func Splitstr(r rune) bool {
	return r == ':'
}

func main(){
	seckill.Backtime()
	http.HandleFunc("/getip", getip)
	http.ListenAndServe(":8080", nil)	
}