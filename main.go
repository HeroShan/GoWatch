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
func main(){
	s:=te()
	fmt.Println("ess",s)
	
}