package leetcode
//https://leetcode-cn.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func RemoveZeroSumSublists(head *ListNode) *ListNode {
	p := &ListNode{}
	p.Next = head
	h := p
	t := h.Next
	for h.Next != nil {
		sum := 0
		for t != nil {
			sum += t.Val
			if sum == 0 {
				break
			}
			t = t.Next
		}
			if sum == 0 {
				t = t.Next
				h.Next = t
			} else {
				h = h.Next
				t = h.Next
			}
	}
	return p.Next
}

//https://leetcode-cn.com/problems/sorted-merge-lcci/
func Merge(A []int, m int, B []int, n int)  {
	AIndex :=m-1
	BIndex := n-1
	idx := len(A)-1
	for AIndex>=0&&BIndex>=0{
		   if A[AIndex]>B[BIndex]{
				 A[idx] = A[AIndex]
			AIndex--
		}else {
			 A[idx] = B[BIndex]
			BIndex--
		}
		idx--
	}
	for BIndex>=0{
		A[idx] = B[BIndex]
		BIndex--
		idx--
	}
}