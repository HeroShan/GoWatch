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
	db.SingularTable(true)
	db.LogMode(true)
	//defer db.Close()
}
func GetAnwser(){
	var a  Astruct
	err := db.First(&a).Error; if err != nil{
		 fmt.Println(err)
	}
	defer db.Close()
}

func (tru Astruct)GetStruct(sid int){
	defer db.Close()
	rows,err := db.Table("astruct").Select("auther, qtitle, qdescribe").Rows(); if err != nil {
		fmt.Println("rows err :",err)
	}
	for rows.Next(){
		err := rows.Scan(&tru.auther,&tru.qtitle,&tru.qdescribe); if err != nil{
			fmt.Println("scan err :",err)
		}
	}
	return
	
}

func (cnt  Content)GetContent(cid int){
	var bb Content
	db.Table("content").Select("question ,anwser ,qe_id ,sort_id ").Scan(&bb)
	fmt.Println(bb)
	db.Raw("SELECT question ,anwser ,qe_id ,sort_id from content").Scan(&bb)
	fmt.Println(bb)
	return

}



