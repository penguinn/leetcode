package index

import (
	"github.com/penguinn/leetcode/common"
)

// 时间复杂度O(n), 空间复杂度O(1)
func removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	result := &common.ListNode{
		Next: head,
	}
	start := result
	end := result
	for i := 0; i <= n; i++ {
		end = end.Next
	}
	for end != nil {
		start = start.Next
		end = end.Next
	}
	start.Next = start.Next.Next
	return result.Next
}
