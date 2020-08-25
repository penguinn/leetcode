package index

import (
	"github.com/penguinn/leetcode/common"
)

func reverseList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	p1 := head
	p2 := head.Next
	// 一定要把这个指针置成nil，不然会循环打印
	p1.Next = nil
	var p3 *common.ListNode
	for p2 != nil {
		p3 = p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}

	return p1
}
