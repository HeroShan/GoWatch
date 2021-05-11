package leetcode

import (
	"log"
)

type node struct {
	Item  string
	Left  *node
	Right *node
}

type bst struct {
	root *node
}

/*
        m
     k     l
  h    i     j
a  b  c  d  e  f


//先序遍历（根左右）：m k h a b i c d l j e f

//中序遍历（左根右）：a h b k c i d m l e j f

//后序遍历（左右根）：a b h c d i k e f j l m

*/

//先序遍历
func (tree *bst) inOrder() {

	var inner func(n *node)
	inner = func(n *node) {

		if n == nil {
			return
		}
		log.Println(n.Item)
		inner(n.Left)
		inner(n.Right)
	}

	inner(tree.root)
}

//中序
func (tree *bst) midOrder() {

	var inner func(n *node)
	inner = func(n *node) {

		if n == nil {
			return
		}
		inner(n.Left)
		log.Println(n.Item)
		inner(n.Right)
	}

	inner(tree.root)
}

//后序
func (tree *bst) lastOrder() {

	var inner func(n *node)
	inner = func(n *node) {

		if n == nil {
			return
		}
		inner(n.Left)
		inner(n.Right)
		log.Println(n.Item)
	}

	inner(tree.root)
}
func (tree *bst) TreeaddBynode(root string, item string) {
	var inner func(n *node, root string, item string)
	inner = func(n *node, root string, item string) {
		if n == nil {
			return
		}
		inner(n.Left, root, item)
		inner(n.Right, root, item)
		if root == n.Item && n.Left == nil {
			add := &node{Item: item}
			n.Left = add
			return
		}
		if root == n.Item && n.Right == nil {
			add := &node{Item: item}
			n.Right = add
			return
		}
	}
	inner(tree.root, root, item)
}

func (tree *node) TreeAddItem(item string) {

	if tree == nil {
		return
	}
	if tree.Left == nil {
		add := &node{Item: item}
		tree.Left = add
		return
	} else if tree.Right == nil {
		add := &node{Item: item}
		tree.Right = add
		return
	} else {
		tree.Left.TreeAddItem(item)
	}

	// var inner func(n *node, item string)
	// inner = func(n *node, item string) {
	// 	if n == nil {
	// 		return
	// 	}

	// 	if n.Left != nil && n.Right == nil {
	// 		add := &node{Item: item}
	// 		n.Right = add
	// 		return

	// 	} else if n.Left == nil && n.Right == nil {
	// 		add := &node{Item: item}
	// 		n.Left = add
	// 		return
	// 	} else {
	// 		inner(n.Left, item)
	// 		inner(n.Right, item)
	// 	}

	// }
	// inner(tree.root, item)
}
