package leetcode

import(
	"testing"
	"fmt"
)

func TestGcdOfStrings(t *testing.T){
	c := GcdOfStrings("ABCABC","ABC")
	fmt.Printf("c:%s \n",c)
}


func TestDefangIPaddr(t *testing.T){
	c := DefangIPaddr("1.1.1.1")
	fmt.Printf("c:%s \n",c)
}


func TestIsPowerOfThree(t *testing.T){
	c:=IsPowerOfThree(0)
	fmt.Println(c)
}

func TestAddList(t *testing.T){
	var l	ListNode
	c:= l.AddList(10)
	fmt.Printf("funcouit:=%#v\n",c)
}