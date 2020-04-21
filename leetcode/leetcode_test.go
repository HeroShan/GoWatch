package leetcode

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T){
	var grid [][]byte
	grid = [][]byte{{1,0,1,1},{0,0,0,0},{1,0,1,1},{0,0,1,0}}
	c := NumIslands(grid)
	fmt.Printf("%d\n",c)
}

func TestCompressString(t *testing.T){
	c:=CompressString("aabcccccaaa")
	fmt.Println(c)
}

func TestRomanToInt(t *testing.T) {
	num := RomanToInt("MMMCCCLX")
	fmt.Println(num)
	num = RomanToInt("MMMCMLXXII")
	fmt.Println(num)
}

func TestScientificToFloat(t *testing.T) {
	c := ScientificToFloat("1.21123E+06")
	fmt.Println(c)
}

func TestGcdOfStrings(t *testing.T) {
	c := GcdOfStrings("ABCABC", "ABC")
	fmt.Printf("c:%s \n", c)
}

func TestDefangIPaddr(t *testing.T) {
	c := DefangIPaddr("1.1.1.1")
	fmt.Printf("c:%s \n", c)
}

func TestIsPowerOfThree(t *testing.T) {
	c := IsPowerOfThree(0)
	fmt.Println(c)
}

func TestAddList(t *testing.T) {
	l := new(ListNode)
	l.Next = nil
	var i int
	for i = 10; i < 15; i++ {
		l.AddHeadList(i)
		fmt.Printf("TestAddList:%#v,  %#v %d\n", l, l.Next, i)
	}
	PrintList(l)
	fmt.Printf("TestAddList:%#v   ,%#v \n", l, l.Next)

}
