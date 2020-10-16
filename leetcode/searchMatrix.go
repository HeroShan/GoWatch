package leetcode
import(
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func inverted(list *ListNode) *ListNode{
	var node *ListNode
    head := list
    for head != nil {
        head.Next, node, head = node, head, head.Next
    }
    return node
}
func printList(listNode *ListNode) {
    for listNode != nil {
        fmt.Println(listNode.Val)
        listNode = listNode.Next
    }
}
func AddTwoNumbers(l1 *ListNode) *ListNode  {
	nl1 := inverted(l1)
	printList(l1)
	printList(nl1)
	return new(ListNode)
}
func SearchMatrix(matrix [][]int, target int) bool {
    var(
		matrixMax int
		matrixMin int  = 0
	)
	if matrix == nil || len(matrix) == 0 {
        return false
	}
	
	for _,array := range matrix{
		matrixMax = len(array)-1
		if array[matrixMin] > target || array[matrixMax] < target {
			continue
		}else{
			for _,value := range array{
				if value > target{
					continue
				}
				if value == target{
					return true
				}

			}
		}
		
	}
	return false

}

func twoSum(nums []int, target int) []int {
    var result []int
    var kMax,cursor int
    kMax = len(nums)
    for k,v := range nums{
        if k == kMax{
            break
        }
        for cursor = k+1;cursor < kMax;cursor++{
            if v+nums[cursor] == target{
            result = append(result,k)
            result = append(result,cursor)
            }
        }
        
    }
    return result
}
