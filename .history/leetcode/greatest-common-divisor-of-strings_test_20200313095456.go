package leetcode

import(
	"testing"
	"fmt"
)

func TestGcdOfStrings(t *testing.T){
	c := GcdOfStrings("abcdabcd","abcdabc")
	fmt.Printf("c:%s \n",c)
}