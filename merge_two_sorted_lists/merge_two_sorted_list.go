/**
 * 21. Merge Two Sorted Lists
 */
package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

// iterative
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var sortedList *ListNode
	head = sortedList
	sortedList = new(ListNode)

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			sortedList.Next = l1
			l1 = l1.Next
		} else {
			sortedList.Next = l2
			l2 = l2.Next
		}
		sortedList = sortedList.Next
	}
	if l1 == nil {
		sortedList.Next = l2
	}
	if l2 == nil {
		sortedList.Next = l1
	}
	return head.Next
}

// recursion
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
