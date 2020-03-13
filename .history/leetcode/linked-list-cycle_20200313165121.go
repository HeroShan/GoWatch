package leetcode


type ListNode struct {
    Val 	int
  	Next 	*ListNode
}

// func hasCycle(head *ListNode) bool {
    
// }

func (this ListNode)CreatList(n int){
		for i:=0;i<n;i++{
			temp := this
			this.Val = i
			temp = this.Next

		}
}