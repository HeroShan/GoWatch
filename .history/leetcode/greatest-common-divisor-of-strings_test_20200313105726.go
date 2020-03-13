package leetcode

import(
	"testing"
	"fmt"
)

func TestGcdOfStrings(t *testing.T){
	c := GcdOfStrings("ABCDEF","ABAB")
	fmt.Printf("c:%s \n",c)
}