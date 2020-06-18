package index

import (
	"github.com/penguinn/leetcode/common"
)

func hasCycle(head *common.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	first := head
	second := head
	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next.Next
		if first == second {
			return true
		}
	}
	return false
}
