package leetcode

import(
	"testing"
	"fmt"
)

func TestGcdOfStrings(t *testing.T){
	c := GcdOfStrings("ABABAB","ABAB")
	fmt.Printf("c:%s \n",c)
}