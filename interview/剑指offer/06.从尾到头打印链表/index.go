package index

type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var p1 *ListNode
	p2 := head
	for p2 != nil {
		p3 := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}

	result := []int{}
	for p1 != nil {
		result = append(result, p1.Val)
		p1 = p1.Next
	}

	return result
}