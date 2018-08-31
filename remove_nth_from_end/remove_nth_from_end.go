/**
 * 19. Remove Nth Node From End of List
**/

package remove_nth_from_end

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	level := 0
	return pRemoveNthFromEnd(head, n, &level)
}

func pRemoveNthFromEnd(head *ListNode, n int, level *int) *ListNode {
	if head == nil {
		*level = 0
		return head
	}
	next := pRemoveNthFromEnd(head.Next, n, level)
	*level++
	if *level == n {
		return next
	}

	head.Next = next
	return head
}
