package leetcode

import(
	"testing"
	"fmt"
)

func TestGcdOfStrings(t *testing.T){
	c := GcdOfStrings("abcdabcd","abcabc")
	fmt.Printf("c:%s \n",c)
}