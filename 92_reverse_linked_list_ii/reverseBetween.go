package reverseBetweenII

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy

	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	then := cur.Next
	for i := 0; i < n-m; i++ {
		cur.Next = then.Next
		then.Next = pre.Next
		pre.Next = then
		then = cur.Next
	}

	return dummy.Next
}
