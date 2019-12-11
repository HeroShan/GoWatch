package qa

import(
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var(
	engine *xorm.Engine
)


func init(){
	var err error
	engine, err = xorm.NewEngine("mysql", "admin:admin@tcp(47.104.225.152:3306)/qasystem?charset=utf8"); if err != nil{
		fmt.Println("connect err:",err)
	}
	engine.ShowSQL(true)
	engine.SetTableMapper(core.SnakeMapper{})
	// engine.Logger().SetLevel(core.LOG_DEBUG)
	//defer engine.Close()
}

func (cnt Content)GetContent(cid int){
	fmt.Println(cnt)
	engine.Where("content_id = 3").Get(&cnt)
	fmt.Println(cnt)
	
}


