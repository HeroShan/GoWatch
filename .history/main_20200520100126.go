package main
import(
	// "fmt"
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
	router.Static("/css/","./css/")
	router.LoadHTMLGlob("./css/*")
	v1 := router.Group("v1")
	{
		v1.GET("/",getip)
		// v1.Get("/logintpl",logintpl)
		// v1.Post("/login",login)
	}
	

	
	router.Run(":8080")
	
}