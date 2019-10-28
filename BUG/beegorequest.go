package controllers

import (
	_"blog/models"
	 "fmt"
	 "encoding/json"
)
//	copyrequestbody = true		app.conf里面设置 	struct里的字段一定要大写
type LoginController struct {
	BaseController
}
type User struct {
	Name     string
	Password string
}

func (c *LoginController) Login() {
	c.TplName = "login/login.html"

}
func (c *LoginController) DoLogin() {
	if c.IsAjax() {
		var u  User
		  err := json.Unmarshal(c.Ctx.Input.RequestBody, &u); if err !=nil {
			fmt.Println("json error QQQ",err) 
		  }
		  fmt.Printf("%T",c.Ctx.Input.RequestBody)
		 pp := &User{ Name:u.Name, Password:u.Password }
		 fmt.Println(u)
		 fmt.Println(pp)
		 
		//models.A()

	}
}
