package leetcode

import(
	"testing"
	"fmt"
)

func TestScientificToFloat(t *testing.T){
	ScientificToFloat("1.21E-03")
}


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
	 l := new(ListNode)
	 l.Next = nil
	var i int
		for i= 10;i<15;i++{
			l.AddHeadList(i)
			fmt.Printf("TestAddList:%#v,  %#v %d\n",  l,l.Next,i)
		}
		PrintList(l)
		fmt.Printf("TestAddList:%#v   ,%#v \n",  l,l.Next)
		
		 
	
	 
	
}