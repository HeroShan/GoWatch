package main
import(
	// "mime"
	"github.com/gin-gonic/gin"
	"GoWatch/mapapi"
	_ "GoWatch/seckill"
	"net/http"
	"strings"
	// "text/template"
	 "GoWatch/auth"
	 "GoWatch/createToken"
	// "time"
	 cl "GoWatch/current_limiter"
	
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

func loginc(c *gin.Context){
		username := c.PostForm("username")
		password := c.PostForm("password")
		Cauth := auth.Check(username,password)
		if Cauth == false {
			c.Redirect(301,"http://"+c.Request.Host+"/login")
		}
		if Cauth == true {
			host := strings.Split(c.Request.Host,":")
			sToken := createToken.GetToken()
			c.SetCookie("wisheart",sToken,7*24*60*60,"/",host[0],false,true)
			c.Redirect(200,"http://"+c.Request.Host+"/admin")
		}
}

func fmsgetip(c *gin.Context){
	//ip := strings.Split(c.Request.RemoteAddr,":")
	area := getiarea("180.101.49.11")
	c.String(200,area)
}

func getiarea(ip string) (path string) {
	var(
		Point mapapi.JsPoint
		areainfo  string
	)
	Point,areainfo = mapapi.Getpoint(ip)
	area := mapapi.Getarea(Point)
	if area != "" {
		path = area
	}else{
		path = areainfo
	}
	return path
}

func LoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		host := strings.Split(c.Request.Host,":")
		if !cl.Serlock(host[0]) {
			c.Redirect(200,"http://"+c.Request.Host+"/error")
		}else{
			cookie,_ := c.Cookie("wisheart")
			if cookie != ""{
				expire := createToken.IsLogin(cookie)
						if expire <=0 {
							sToken := createToken.GetToken()
							c.SetCookie("wisheart",sToken,0,"/",host[0],false,true)
							c.Redirect(200,"http://"+c.Request.Host+"/login")
						}else{
							c.Next()
						}
			}else{
				c.Redirect(200,"http://"+c.Request.Host+"/login")
			}
		}

	}

}

func error(c *gin.Context){
	c.HTML(http.StatusOK,"error.html",gin.H{})
}

func adminPage(c *gin.Context){
	c.HTML(http.StatusOK,"admin.html",gin.H{})
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
	login := router.Group("")
	{
		login.GET("/",getip)
		login.GET("/login",logintpl)
		login.POST("/login",loginc)
		login.GET("/fmsgetip",fmsgetip)
		login.GET("/error",error)
		// v1.Post("/login",login)
	}

	admin := router.Group("admin")
	{
		admin.Use(LoginMiddleware())
		admin.GET("/",adminPage)

	}

	

	
	router.Run(":8080")
	
}