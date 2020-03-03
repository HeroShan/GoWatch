package oslearn

import(
	"testing"
	"fmt"
)


func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	fmt.Printf("%#v \n",res)
}