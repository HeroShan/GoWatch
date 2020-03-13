package leetcode


type ListNode struct {
    Val 	int
  	Next 	*ListNode
}

// func hasCycle(head *ListNode) bool {
    
// }

func (this ListNode)CreatList(n int){
		temp := this
		for i:=0;i<n;i++{
			temp = &temp.Next
			this.Val = i
			this.Next = &temp

		}
}