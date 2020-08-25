package index

import (
	"github.com/penguinn/leetcode/common"
)

func getIntersectionNode(headA, headB *common.ListNode) *common.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p := headA
	q := headB

	pFlag := 0
	qFlag := 0
	for p != q {
		p = p.Next
		q = q.Next
		if p == nil && pFlag != 1 {
			p = headB
			pFlag++
		} else if p == nil && pFlag == 1 {
			break
		}
		if q == nil && qFlag != 1 {
			q = headA
			qFlag++
		} else if q == nil && qFlag == 1 {
			break
		}
	}

	if p == q {
		return p
	} else {
		return nil
	}
}
