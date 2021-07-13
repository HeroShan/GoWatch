package leetcode

import (
	"fmt"
	"testing"
)

func TestBNode_EachBn(t *testing.T) {
	bt := NewBtree()
	bt.AddBtree(nil, "1")
	bt.AddBtree(nil, "2")
	bt.EachBt()
}

func TestBTree_AddBtree(t *testing.T) {
	bt := NewBtree()
	bt.AddBtree(nil, "1")
	fmt.Printf("%#v\n", bt.treeList[0])
}

func TestNewBtree(t *testing.T) {
	b := NewBtree()
	fmt.Printf("%#v\n", b)
}
