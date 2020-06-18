package index

import (
	"github.com/penguinn/leetcode/common"
)

// 时间复杂度O(m+n), 空间复杂度O(1)
func mergeTwoLists(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	result := new(common.ListNode)
	p := result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next

	}
	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}
	return result.Next
}
