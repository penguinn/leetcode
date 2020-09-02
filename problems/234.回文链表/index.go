package index

import (
	"github.com/penguinn/leetcode/common"
)

func isPalindrome(head *common.ListNode) bool {
	if head == nil {
		return true
	}
	if head.Next == nil {
		return true
	}
	i := head
	j := head
	for j != nil && j.Next != nil {
		i = i.Next
		j = j.Next
		if j == nil {
			break
		} else {
			j = j.Next
		}
	}

	p1 := i
	p2 := i.Next
	p1.Next = nil
	var p3 *common.ListNode
	for p2 != nil {
		p3 = p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}

	for head != nil && p1 != nil {
		if head.Val != p1.Val {
			return false
		}
		head = head.Next
		p1 = p1.Next
	}

	return true
}
