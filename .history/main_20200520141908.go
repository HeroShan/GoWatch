package main
import(
	 _"fmt"
	// "mime"
	"github.com/gin-gonic/gin"
	//"GoWatch/mapapi"
	_ "GoWatch/seckill"
	"net/http"
	"strings"
	// "text/template"
	 "GoWatch/auth"
	 "GoWatch/createToken"
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
		password := c.PostForm("password")
		Cauth := auth.Check(username,password)
		if Cauth == false {
			c.Redirect(http.StatusMovedPermanently,c.Request.Host+"/login")
		}
		if Cauth == true {
			host := strings.Split(c.Request.Host,":")
			sToken := createToken.GetToken()
			c.SetCookie("wisheart",sToken,7*24*60*60,"/",host[0],false,true)
			
		}
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