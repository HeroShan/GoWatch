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
func (bt *BTree) Truncate() {
	*bt = BTree{nil, nil}
	return
}

//遍历btree  所有元素
func (bt *BTree) EachBt() {
	for bk, bn := range bt.treeList {
		i := 0
		bn.EachBn(bk, i)
	}
}

//遍历每个枝上的节点
func (bn *bNode) EachBn(bk int, nk int) {
	if bn.next == nil {
		return
	}
	fmt.Println("limb is:", bk, "node is:", nk, "data is:", bn.data)
	bn = bn.next
	nk++
	bn.EachBn(bk, nk)
}

//根据头进行追加btree 元素
//头为空  追加到根
//头不为空  找到头追加到下一个节点   next顺移到下一个节点
func (bt *BTree) AddBtree(head interface{}, item interface{}) bool {
	var nilHead, next bNode
	if head == nil {
		nilHead.data = item
		nilHead.next = &next
		bt.treeList = append(bt.treeList, &nilHead)
		return true
	} else {
		for _, v := range bt.treeList {
			for v.next != nil {
				if v.data == head {
					if v.next == nil {
						nilHead.data = item
						nilHead.next = &next
						v.next = &nilHead
						return true
					} else {
						next = *v.next
						nilHead.data = item
						nilHead.next = &next
						v.next = &nilHead
						return true
					}
				}
				v = v.next
			}

		}
	}
	return false
}
