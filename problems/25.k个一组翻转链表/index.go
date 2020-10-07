package index

import (
	. "github.com/penguinn/leetcode/common"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	result := &ListNode{Next: head}
	pre := result
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return result.Next
			}
		}

		next := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}

	return result.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	p1 := tail.Next
	p2 := head
	for p1 != tail {
		p3 := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}

	return tail, head
}
