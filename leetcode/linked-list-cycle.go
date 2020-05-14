package leetcode
import(
	"fmt"
)
type ListNode struct {
    Val 	int
  	Next 	*ListNode
}

func (this *ListNode)AddTailList(n int) {
	var temp ListNode
	
		for this.Next != nil{
			this = this.Next
		}
		temp.Val = n
		temp.Next = nil
		this.Next = &temp	
			
}

func (this *ListNode)AddHeadList(n int) {
	var temp ListNode
	temp.Next = nil
	temp.Val = n
	for{
		if this.Next ==nil {
			this.Next = &temp
			break
		}else{
			this = this.Next
		}
	}
	
}


func PrintList(list *ListNode) {
	p := list
	for{
		
		if p.Next != nil {
			fmt.Println(p.Val)
			p = p.Next
		}else{
			fmt.Println(p.Val)
			break
		}
		
		
	}
}