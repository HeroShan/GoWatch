package qa

import(
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strconv"
	"strings"
)

var(
	engine *xorm.Engine
)


func init(){
	var err error
	engine, err = xorm.NewEngine("mysql", "admin:admin@tcp(47.104.225.152:3306)/qasystem?charset=utf8"); if err != nil{
		fmt.Println("connect err:",err)
	}
	engine.ShowSQL(true)	//开启sql调试
	engine.SetTableMapper(core.SnakeMapper{})
	//engine.Logger().SetLevel(core.LOG_DEBUG)
	//defer engine.Close()
}


func (tru *Astruct)Getqa(cid int) *[]ContentAnwser{
	var con []ContentAnwser
	engine.Where("Struct_id = "+strconv.Itoa(cid)).Get(tru)
	fmt.Println(tru.Content_id)
	engine.Table("anwser").Join("INNER","content","content.content_id = anwser.content_id").In("content.content_id",strings.Split(tru.Content_id,",")).Find(&con)
	return &con 
	// if err != nil{
	// 	fmt.Println(err)
	// }
	// for rows.Next(){
	// 	rows.Scan(&con)
	// 	fmt.Println(&con)
	// }
	
	
}


