package main

import(
	"GoWatch/orm/qa"
	_"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)



func main(){
	cnt := new(qa.Astruct)
	qa := cnt.Getqa(15)
	// tru.GetContent(15)
	fmt.Println(*qa)
}