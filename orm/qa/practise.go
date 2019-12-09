package qa

import(
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var(
	db *gorm.DB
)


func init(){
	db, _ = gorm.Open("mysql", "admin:admin@(47.104.225.152:3306)/qasystem?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	db.SingularTable(true)
	//defer db.Close()
}
func GetAnwser(){
	var a  Astruct
	err := db.First(&a).Error; if err != nil{
		 fmt.Println(err)
	}
	defer db.Close()
}

func (tru *Astruct)GetContent(sid int){
	tru.struct_id = sid
	err := db.Table("astruct").Select("auther, title, `describe`, content_id, date, counts").Where("struct_id=? ",tru.struct_id).Scan(&tru).GetErrors(); if err != nil{
		fmt.Println("db find-",err)
	}
	fmt.Println(tru)
	defer db.Close()
}


