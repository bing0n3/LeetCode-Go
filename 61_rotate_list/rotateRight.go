package rotatelist

/**
 * 61.Rotate List
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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	tail := head
	length := 1
	for tail.Next != nil {
		tail = tail.Next
		length++
	}
	tail.Next = head

	k = k % length

	if k != 0 {
		for i := 0; i < length-k; i++ {
			tail = tail.Next
		}
	}
	newH := tail.Next
	tail.Next = nil
	return newH
}
