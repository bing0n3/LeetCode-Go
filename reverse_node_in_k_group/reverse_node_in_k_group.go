/**
 * 25. Reverse Nodes in k-Group
 */

package reverse_node_in_k_group

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	pre := new(ListNode)
	pre.Next = head
	curr := pre
	swap := make([]*ListNode, k)
	for i := 0; i < k; i++ {
		swap[i] = new(ListNode)
	}
	pos := curr
	for {
		pos = curr
		for i := 0; i < k; i++ {
			if curr.Next != nil {
				swap[i] = curr.Next
				curr = curr.Next
			} else {
				return pre.Next
			}
		}
		curr = pos
		swap[0].Next = swap[k-1].Next
		for i := k - 1; i >= 0; i-- {
			curr.Next = swap[i]
			curr = curr.Next
		}
	}
}
