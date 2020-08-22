package index

import (
	"github.com/penguinn/leetcode/common"
)

// 还可以通过栈的思维来解决
func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	ptr1 := l1
	ptr2 := l1.Next
	ptr1.Next = nil
	for ptr2 != nil {
		ptr3 := ptr2.Next
		ptr2.Next = ptr1
		ptr1 = ptr2
		ptr2 = ptr3
	}
	l1 = ptr1

	ptr1 = l2
	ptr2 = l2.Next
	ptr1.Next = nil
	for ptr2 != nil {
		ptr3 := ptr2.Next
		ptr2.Next = ptr1
		ptr1 = ptr2
		ptr2 = ptr3
	}
	l2 = ptr1

	flag := 0
	result := l1
	ptr := l1
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + flag
		flag = val / 10
		l1.Val = val % 10
		ptr = l1
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		val := l1.Val + flag
		flag = val / 10
		l1.Val = val % 10
		ptr = l1
		l1 = l1.Next
	}

	for l2 != nil {
		val := l2.Val + flag
		flag = val / 10
		ptr.Next = &common.ListNode{Val: val % 10}
		ptr = ptr.Next
		l2 = l2.Next
	}

	if flag != 0 {
		ptr.Next = &common.ListNode{Val: flag}
		ptr = ptr.Next
	}

	ptr1 = result
	ptr2 = result.Next
	ptr1.Next = nil
	for ptr2 != nil {
		ptr3 := ptr2.Next
		ptr2.Next = ptr1
		ptr1 = ptr2
		ptr2 = ptr3
	}

	return ptr1
}
