package leetcode
import(
	"fmt"
)

type ListNode struct {
    Val 	int
  	Next 	*ListNode
}

// func hasCycle(head *ListNode) bool {
    
// }

func (this ListNode)AddList(n int) ListNode {
		temp := this
		temp.Val = n
		fmt.Printf("temp: = %#v\n",temp)
		temp.Next = &this
		return temp 

}