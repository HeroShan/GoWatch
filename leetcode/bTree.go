package leetcode

import "fmt"

type BTree struct {
	treeList []*bNode
	next     *BTree
}

type bNode struct {
	data interface{}
	next *bNode
}

func NewBtree() *BTree {
	return new(BTree)
}

func (bt *BTree) EachBt() {
	for bk, bn := range bt.treeList {
		i := 0
		bn.EachBn(bk, i)
	}
}
func (bn *bNode) EachBn(bk int, nk int) {
	if bn.next == nil {
		return
	}
	fmt.Println("limb is:", bk, "node is:", nk, "data is:", bn.data)
	bn = bn.next
	nk++
	bn.EachBn(bk, nk)
}

func (bt *BTree) AddBtree(head interface{}, item interface{}) {
	if head == nil {
		var nilHead, next bNode
		nilHead.data = item
		nilHead.next = &next
		bt.treeList = append(bt.treeList, &nilHead)
		return
	} else {

	}

}
