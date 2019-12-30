package main

import(
	"GoWatch/orm/qa"
	_ "github.com/go-sql-driver/mysql"
	"encoding/hex"
	"fmt"
)



func main(){
	
	cnt := new(qa.User)
	cnt.Username = hex.EncodeToString([]byte("※"))
	cnt.Password = "123123"
	cnt.Lastime  = 1020202
	cc := cnt.Add()
	fmt.Println(cc)
	// tru.GetContent(15)
}