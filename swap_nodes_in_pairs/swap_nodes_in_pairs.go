// 24. Swap Nodes in Pairs

package swap_nodes_in_pairs

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	pre := new(ListNode)
	pre.Next = head
	curr := pre
	a := new(ListNode)
	b := new(ListNode)

	for curr.Next != nil && curr.Next.Next != nil {
		a = curr.Next
		b = curr.Next.Next
		a.Next = b.Next
		curr.Next = b
		curr.Next.Next = a
		curr = curr.Next.Next
	}

	return pre
}
