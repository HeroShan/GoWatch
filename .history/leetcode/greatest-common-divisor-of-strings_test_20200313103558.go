package leetcode

import(
	"testing"
	"fmt"
)

func TestGcdOfStrings(t *testing.T){
	c := GcdOfStrings("abcdabcd","abcde")
	fmt.Printf("c:%s \n",c)
}