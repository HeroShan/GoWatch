package main
import(
	 "github.com/juju/errors"
	 "GoWatch/getuuid"
	e "errors"
	"fmt"
)

func te() string{
	var i int
	var cr string
	i = i+7
	if i>3{
	ter :=	e.New("number out to range")
	fmt.Println(ter)
	cr = errors.ErrorStack(ter)
	}
	return cr
	
	
}
func uuid(){
    ui :=	getuuid.GetV4()
	fmt.Println(ui)
}

var snum int
func main(){
	uuid()
	uuid()
	uuid()
	//s:=te()
	// fmt.Println("ess",s)
	// for i := 0; i<=10;i++{
	// 	snum = i
	// 	switch snum {
	// 	case 1,2,3 :
	// 		fmt.Println(1)
		
	// 	case 4,5,6 :
	// 		fmt.Println(2)
		
	// 	case 7,8,9 :
	// 		fmt.Println(3)
	// 	}
	// }  
	

}