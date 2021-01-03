package index

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 把链表扩为原来2倍
	p := head
	for p != nil {
		addNode := &Node{Val: p.Val, Next: p.Next}
		p.Next = addNode
		p = p.Next.Next
	}

	// 为复制的链表填入random指针
	p = head
	for p != nil {
		if p.Random == nil {
			p.Next.Random = nil
		} else {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	// 拆分出复制的链表
	hair := &Node{}
	q := hair
	p = head
	for p != nil {
		q.Next = p.Next
		q = q.Next
		p.Next = p.Next.Next
		p = p.Next
	}

	return hair.Next
}
