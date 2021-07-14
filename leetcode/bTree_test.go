package leetcode

import (
	"fmt"
	"testing"
)

func TestBNode_EachBn(t *testing.T) {
	bt := NewBtree()
	bt.AddBtree(nil, "1")
	bt.AddBtree(nil, "5")
	bt.EachBt()
}

func TestBTree_AddBtree(t *testing.T) {

}

func TestNewBtree(t *testing.T) {
	b := NewBtree()
	fmt.Printf("%#v\n", b)
}
