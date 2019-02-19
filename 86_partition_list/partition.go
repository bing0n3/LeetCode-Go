package partitionList

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	lessList := new(ListNode)
	geList := new(ListNode)
	dummyL := lessList
	dummyGe := geList

	for head != nil {
		if head.Val < x {
			lessList.Next = head
			lessList = lessList.Next
		} else {
			geList.Next = head
			geList = geList.Next
		}
		head = head.Next
	}
	geList.Next = nil // avoid cycle link
	lessList.Next = dummyGe.Next

	return dummyL.Next
}
