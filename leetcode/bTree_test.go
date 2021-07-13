package leetcode

import (
	"fmt"
	"testing"
)

func TestBTree_Truncate(t *testing.T) {
	var r bool
	bt := new(BTree)
	r = bt.AddBtree(nil, 0)
	fmt.Println(r)
	for i := 0; i < 2; i++ {
		p := i + 1
		r = bt.AddBtree(i, p)
		fmt.Println(r, i, p)
	}

	bt.AddBtree(nil, 11)
	for j := 11; j < 30; j++ {
		bt.AddBtree(j, j+1)
	}
	bt.EachBt()
	bt.Truncate()
	fmt.Println("truncate!")
	bt.EachBt()
}

func BenchmarkBNode_EachBn(b *testing.B) {
	var r bool
	bt := new(BTree)
	r = bt.AddBtree(nil, 0)
	fmt.Println(r)
	for i := 0; i < 2; i++ {
		p := i + 1
		r = bt.AddBtree(i, p)
		fmt.Println(r, i, p)
	}

	bt.AddBtree(nil, 11)
	for j := 11; j < 30; j++ {
		bt.AddBtree(j, j+1)
	}
	bt.EachBt()
	for i := 0; i < b.N; i++ { //use b.N for looping
		bt.EachBt()
	}
}

func TestBNode_EachBn(t *testing.T) {
	bt := NewBtree()
	bt.AddBtree(nil, "1")
	bt.AddBtree(nil, "5")
	bt.EachBt()
}

func TestBTree_AddBtree(t *testing.T) {
	var r bool
	bt := NewBtree()
	r = bt.AddBtree(nil, 0)
	fmt.Println(r)
	for i := 0; i < 2; i++ {
		p := i + 1
		r = bt.AddBtree(i, p)
		fmt.Println(r, i, p)
	}

	bt.AddBtree(nil, 11)
	for j := 11; j < 30; j++ {
		bt.AddBtree(j, j+1)
	}
	bt.EachBt()

}

func TestNewBtree(t *testing.T) {
	b := NewBtree()
	fmt.Printf("%#v\n", b)
}
