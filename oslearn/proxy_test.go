package oslearn

import(
	"testing"
	"fmt"
)


func TestProxy(t *testing.T){
	var sub Subject
	fmt.Printf("1%#v\n ",sub)
	sub = &Proxy{}
	fmt.Printf("2%#v\n ",sub)
	res := sub.Do()
	fmt.Printf("3%#v\n %#v\n",sub,res)
}