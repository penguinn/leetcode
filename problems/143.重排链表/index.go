package index

import (
	"github.com/penguinn/leetcode/common"
)

func reorderList(head *common.ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 使用快慢指针找出中间节点
	slow := head
	quick := head
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}

	// 翻转链表
	var p1 *common.ListNode
	p2 := slow
	var p3 *common.ListNode
	for p2 != nil {
		p3 = p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}

	front := head
	backend := p1
	current := &common.ListNode{}
	i := 0
	for current != slow {
		if i%2 == 1 {
			current.Next = backend
			current = current.Next
			backend = backend.Next
		} else {
			current.Next = front
			current = current.Next
			front = front.Next
		}
		i++
	}
}
