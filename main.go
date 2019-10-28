package main
import(
	 "github.com/juju/errors"
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
var snum int
func main(){
	s:=te()
	fmt.Println("ess",s)
	for i := 0; i<=6;i++{
		snum = i
		switch snum {
		case 1,2,3 :
			fmt.Println(1)
		
		case 4,5,6 :
			fmt.Println(2)
		}
	}  
	

}