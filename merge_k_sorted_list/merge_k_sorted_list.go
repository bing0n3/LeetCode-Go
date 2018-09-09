/**
 * 23. Merge k Sorted Lists
 */
package merge_k_sorted_list

import (
	"fmt"
	"sort"
)

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	nodes := make([]int, 0)
	head := new(ListNode)
	current := head

	for _, l := range lists {
		for l != nil {
			nodes = append(nodes, l.Val)
			l = l.Next
		}
	}
	sort.Ints(nodes)
	for _, x := range nodes {
		current.Next = new(ListNode)
		current.Next.Val = x
		current = current.Next
	}

	return head.Next
}

type minHeap []*ListNode

func mergeKLists2(lists []*ListNode) *ListNode {
	head := new(ListNode)
	current := head
	for {
		min := -1
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if min == -1 {
				min = i
			}
			if lists[i].Val < lists[min].Val {
				min = i
			}
		}
		if min == -1 {
			fmt.Println(head.Next.Val)
			return head.Next
		}
		current.Next = lists[min]
		lists[min] = lists[min].Next
		current = current.Next
	}
}
