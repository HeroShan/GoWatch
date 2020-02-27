package oslearn

import(
	"testing"
	"fmt"
)


func TestGetInstance(t *testing.T){
	var Single Singleton

	Single.Name = "F"
	Single.Age 	= 25

	c := GetInstance(Single)
	
	fmt.Printf("%#v \n",c)
}