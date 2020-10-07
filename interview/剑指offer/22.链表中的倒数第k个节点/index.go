package index

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	p := head
	for i:= 1;i<=k;i++ {
		p = p.Next
		if p == nil {
			return head
		}
	}

	for p != nil {
		p = p.Next
		head = head.Next
	}

	return head
}
