package deleteDuplicates_ii

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

func deleteDuplicates(head *ListNode) *ListNode {
	prv := &ListNode{Next: head}
	tmp := prv
	cur := head

	for cur != nil {
		// the next is nil means, we can decide to delete or not now
		// when not duplicate, we should pointer prv.Next
		if cur.Next == nil || cur.Val != cur.Next.Val {
			// if no duplicate
			if prv.Next == cur {
				prv = cur
			} else {
				// if duplicate between prv and cur
				prv.Next = cur.Next
			}
		}
		// not matter duplicate or not, move position
		cur = cur.Next
	}

	return tmp.Next
}
