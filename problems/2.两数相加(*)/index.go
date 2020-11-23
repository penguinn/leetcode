package index

import (
	"github.com/penguinn/leetcode/common"
)

// 直接让链表从头到尾开始相加，一定要注意好进位的问题
func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	addBit := 0
	head := &common.ListNode{}
	p := head
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + addBit
		addBit = val / 10
		l1.Val = val % 10
		p.Next = l1
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		val := l1.Val + addBit
		addBit = val / 10
		l1.Val = val % 10
		p.Next = l1
		p = p.Next
		l1 = l1.Next
	}

	for l2 != nil {
		val := l2.Val + addBit
		addBit = val / 10
		l2.Val = val % 10
		p.Next = l2
		p = p.Next
		l2 = l2.Next
	}

	if addBit == 1 {
		p.Next = &common.ListNode{Val: 1}
	}

	return head.Next
}
