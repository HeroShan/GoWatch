package main
import(
	 "fmt"
	// "mime"
	"github.com/gin-gonic/gin"
	//"GoWatch/mapapi"
	_ "GoWatch/seckill"
	"net/http"
	// "strings"
	// "text/template"
	// "GoWatch/auth"
	// "GoWatch/createToken"
	// "time"
	// cl "GoWatch/current_limiter"
	
)



 
type Iperr struct{
	Info string 
}


func getip(c *gin.Context){
		c.HTML(http.StatusOK,"index.html",gin.H{
			"title" : "King",
		})
}

func logintpl(c *gin.Context){
		c.HTML(http.StatusOK,"login.html",gin.H{})
}

func login(c *gin.Context){
		username := c.PostForm("username")
		password := c.password("username")
		fmt.Println(username)
		fmt.Println(password)
		// Cauth := auth.Check(username[0],password[0])
		// if Cauth == false {
		// 	w.Write([]byte("<script>alert('登陆失败'); window.history.back(-1); </script>"))	 
		// }
		// if Cauth == true {
		// 	sToken := createToken.GetToken()
		// 	sCookie := &http.Cookie{
		// 		Name:   "wisheart",
		// 		Value:  sToken,
		// 		Path:   "/",
		// 		Domain: r.Host ,
		// 		MaxAge: 7*24*60*60,
		// 	}
		// 	http.SetCookie(w, sCookie)
		// 	w.Write([]byte("<script>self.location.href='/admin'</script>"))
		// }
}

func main(){
	// go monitoring()
	// mime.AddExtensionType(".js", "text/javascript")	//static
	// http.HandleFunc("/", getip)
	// http.HandleFunc("/fmsgetip", fmsgetip)
	// http.HandleFunc("/login", login)
	// http.HandleFunc("/error", errorfn)
	// http.Handle("/admin", LoginMiddleware(http.HandlerFunc(admin)))	//	登录
	// http.Handle("/admin/simpleUpload", LoginMiddleware(http.HandlerFunc(simpleUpload)))	//	单文件上传
	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css")))) //static
	// http.ListenAndServe(":8080", nil)
	router := gin.Default()
	router.Static("/css","./css")
	router.LoadHTMLGlob("css/html/*")
	v1 := router.Group("")
	{
		v1.GET("/",getip)
		v1.GET("/login",logintpl)
		v1.POST("login",login)
		// v1.Post("/login",login)
	}
	

	
	router.Run(":8080")
	
}