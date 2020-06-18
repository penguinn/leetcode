package index

import (
	"github.com/penguinn/leetcode/common"
)

func reverseList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var first *common.ListNode
	second := head
	third := head.Next
	for third != nil {
		second.Next = first
		first = second
		second = third
		third = third.Next
	}
	second.Next = first

	return second
}
