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