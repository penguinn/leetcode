package index

import (
	"github.com/penguinn/leetcode/common"
)


// 直接让链表从头到尾开始相加，一定要注意好进位的问题
func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	addNum := 0
	result := l1
	var current *common.ListNode
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + addNum
		addNum = val / 10
		l1.Val = val % 10
		current = l1
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		val := l1.Val + addNum
		addNum = val / 10
		l1.Val = val % 10
		current = l1
		l1 = l1.Next
	}

	for l2 != nil {
		val := l2.Val + addNum
		addNum = val / 10
		current.Next = &common.ListNode{Val: val % 10}
		current = current.Next
		l2 = l2.Next
	}

	if addNum != 0 {
		current.Next = &common.ListNode{Val: addNum}
		addNum = 0
	}

	return result
}
