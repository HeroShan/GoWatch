package leetcode

type bTree struct {
	treeList []*bNode
	next     *bTree
}

type bNode struct {
	data interface{}
	next *bNode
}

func NewBtree() *bTree {
	return new(bTree)
}

//清除整个b树
func (bt *bTree) Truncate() {
	*bt = bTree{nil, nil}
	return
}

//遍历btree  所有元素
func (bt *bTree) EachBt() {
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
	//fmt.Println("limb is:", bk, "node is:", nk, "data is:", bn.data)
	bn = bn.next
	nk++
	bn.EachBn(bk, nk)
}

//is  funny
//在头部确定数据链的区间前后差值是  前值的二倍 区间插入
func (bt *bTree) MidInsertBtree(item int) {

	lenBtList := len(bt.treeList)
	var next bNode
	if lenBtList == 0 || lenBtList == 1 {
		next.data = item
		next.next = nil
		bt.treeList = append(bt.treeList, &next)
	} else {
		if item > 2*bt.treeList[1].data.(int) {

		}
	}
	return
}

//根据头进行追加b+tree 元素
//头为空  追加到根
//头不为空  找到头追加到下一个节点   next顺移到下一个节点
func (bt *bTree) AddBtree(head interface{}, item interface{}) bool {
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
