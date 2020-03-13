package leetcode


type ListNode struct {
    Val 	int
  	Next 	*ListNode
}

// func hasCycle(head *ListNode) bool {
    
// }

func (this ListNode)AddList(n int) ListNode {
		temp := this
		temp.Val = n
		temp.Next = &this
		this.Next = temp
		return this 

}