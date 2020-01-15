package main
import(
	"github.com/juju/errors"
	e "errors"
	"fmt"
	"mime"
	"GoWatch/mapapi"
	_ "GoWatch/seckill"
	"net/http"
	"strings"
	"text/template"
	"GoWatch/auth"
	"GoWatch/createToken"
	"time"
	cl "GoWatch/current_limiter"
)
//	practiced Error Package
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



var(
	snum int
	ip string
	Point mapapi.JsPoint
	areainfo string
	path string
) 
 
type Iperr struct{
	Info string 
}
// get IP interface
func getip(w http.ResponseWriter, r *http.Request){
	ipstring := r.RemoteAddr
	ip1 := strings.FieldsFunc(ipstring,Splitstr)
	ip = ip1[0]
	//ip = "180.101.49.11"
	//ip = "183.131.107.149"
	Point,areainfo = mapapi.Getpoint(ip)
	
	//Point.X = "32.05723550"
	//Point.Y = "120.21287663"
	area := mapapi.Getarea(Point)
	if area != "" {
		path = area
	}else{
		path = areainfo
	}
	html, err := template.ParseFiles("./css/index.html")
	if err != nil{
		fmt.Println(err)
		}
	html.Execute(w, path)
	mapapi.Connect(ip,path,Point)
}

//split string function
func Splitstr(r rune) bool {
	return r == ':'
}

//User Login function
func login(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		html, _ := template.ParseFiles("./css/login.html")
		html.Execute(w, nil)
	}
	if r.Method == "POST"{
		r.ParseForm()
		username := r.PostForm["username"]
		password := r.PostForm["password"]
		Cauth := auth.Check(username[0],password[0])
		if Cauth == false {
			w.Write([]byte("<script>alert('登陆失败'); window.history.back(-1); </script>"))	 
		}
		if Cauth == true {
			sToken := createToken.GetToken()
			sCookie := &http.Cookie{
				Name:   "wisheart",
				Value:  sToken,
				Path:   "/",
				Domain: r.Host ,
				MaxAge: 7*24*60*60,
			}
			http.SetCookie(w, sCookie)
			w.Write([]byte("<script>self.location.href='/admin'</script>"))
		}
	}
}


func LoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		ipstr := strings.Split(r.RemoteAddr,":")
		if !cl.Serlock(ipstr[0]) {
			w.Write([]byte("<script>self.location.href='/error'</script>"))
		}else{
			if r.Method == "GET" {
				sCookie,_ := r.Cookie("wisheart")
				if sCookie == nil{
					http.Redirect(w, r, "/login", http.StatusOK)
				}else{
					if sCookie.Value != ""  {
						expire := createToken.IsLogin(sCookie.Value)
						if expire <=0 {
							sCookie := &http.Cookie{
								Name:   "wisheart",
								Value:  "",
								Path:   "/",
								Domain: r.Host ,
								MaxAge: 0,
							}
							http.SetCookie(w, sCookie)
							http.Redirect(w, r, "/login", http.StatusOK)
						}else{
							next.ServeHTTP(w, r)
						}
					}else{
						http.Redirect(w, r, "/login", http.StatusOK)
					}
				}
			}
		}
			
	})
	
}

func errorfn(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		html, _ := template.ParseFiles("./css/error.html")
		html.Execute(w, nil)
	}
}

func admin(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		html, _ := template.ParseFiles("./css/admin.html")
		html.Execute(w, nil)
	}
}

func simpleUpload(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		html, _ := template.ParseFiles("./css/simpleUpload.html")
		html.Execute(w, nil)
	}
	if r.Method == "POST"{

	}
}

func fmsgetip(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		ipstring := r.RemoteAddr
		w.Write([]byte(ipstring))
	}
}

func monitoring(){
	for{
		time.Sleep(24 * time.Hour)
		cl.SerUnlock()
		wk := time.Now().Weekday().String()
		if wk == "Sunday" || wk == "Wednesday"{
			createToken.DelExpireToken()
		}
	}
}

func main(){
	go monitoring()
	mime.AddExtensionType(".js", "text/javascript")	//static
	http.HandleFunc("/", getip)
	http.HandleFunc("/fmsgetip", fmsgetip)
	http.HandleFunc("/login", login)
	http.HandleFunc("/error", errorfn)
	http.Handle("/admin", LoginMiddleware(http.HandlerFunc(admin)))	//	登录
	http.Handle("/admin/simpleUpload", LoginMiddleware(http.HandlerFunc(simpleUpload)))	//	单文件上传
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css")))) //static
	http.ListenAndServe(":8080", nil)	
}