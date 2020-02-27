package oslearn

import(
	"testing"
	"fmt"
)


func TestGetInstance(t *testing.T){


	c := GetInstance()
	c.Name = "F"
	c.Age  = 25
	b := GetInstance()
	b.Name = "S"
	fmt.Printf("%#v \n",c)
	fmt.Printf("%#v \n",b)
}