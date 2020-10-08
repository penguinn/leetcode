package index

import (
	. "github.com/penguinn/leetcode/common"
)

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	p := &ListNode{Next: head}
	result := p
	if n <= m {
		return p.Next
	}
	for i := 1; i < m; i++ {
		p = p.Next
		if p == nil {
			return result.Next
		}
	}
	prev := p
	p = p.Next
	if p == nil {
		return result.Next
	}
	start := p
	for i := 0; i < n-m; i++ {
		p = p.Next
		if p == nil {
			return result.Next
		}
	}
	p1 := p.Next
	p2 := start
	for p1 != p {
		p3 := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}
	prev.Next = p

	return result.Next
}
