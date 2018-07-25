/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := new(ListNode)
	curr := head
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			l1, n1 = l1.Next, l1.Val
		}
		if l2 != nil {
			l2, n2 = l2.Next, l2.Val
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10, Next: nil}
		curr = curr.Next
		if carry > 0 {
			curr.Next = &ListNode{Val: 1, Next: nil}
		}
	}
	return head.Next
}
