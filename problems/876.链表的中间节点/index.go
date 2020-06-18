package index

import (
	"github.com/penguinn/leetcode/common"
)

func middleNode(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := &common.ListNode{
		Next: head,
	}
	slow := result
	fast := result
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow.Next
}
