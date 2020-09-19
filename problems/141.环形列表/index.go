package index

import (
	"github.com/penguinn/leetcode/common"
)

func hasCycle(head *common.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast && fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		} else {
			break
		}
	}
	
	if fast != nil {
		return true
	} else {
		return false
	}
}