package leetcode

import(
	"testing"
	"fmt"
)

func TestGcdOfStrings(t *testing.T){
	c := GcdOfStrings("abababab","abab")
	fmt.Printf("c:%s \n",c)
}