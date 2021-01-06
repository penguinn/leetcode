package index

import . "github.com/penguinn/leetcode/common"

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	q := head
	p := head.Next
	for p != nil {
		if p.Val == val {
			q.Next = p.Next
			break
		}
		q = p
		p = p.Next
	}

	return head
}
