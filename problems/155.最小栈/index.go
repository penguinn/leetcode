package main

import (
	"fmt"
)

// 时间复杂度为O(1)，空间复杂度为O(n)
type MinStack struct {
	Nodes    []int
	MinNodes []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Nodes:    []int{},
		MinNodes: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.Nodes = append(this.Nodes, x)
	if len(this.MinNodes) == 0 || this.MinNodes[len(this.MinNodes)-1] >= x {
		this.MinNodes = append(this.MinNodes, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.Nodes) > 0 {
		if this.MinNodes[len(this.MinNodes)-1] == this.Nodes[len(this.Nodes)-1] {
			this.MinNodes = this.MinNodes[:len(this.MinNodes)-1]
		}
		this.Nodes = this.Nodes[:len(this.Nodes)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.Nodes) > 0 {
		return this.Nodes[len(this.Nodes)-1]
	} else {
		return 0
	}
}

func (this *MinStack) GetMin() int {
	if len(this.MinNodes) > 0 {
		return this.MinNodes[len(this.MinNodes)-1]
	} else {
		return 0
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // 返回 -3.
	minStack.Pop()
	fmt.Println(minStack.Top()) // 返回 0.
	fmt.Println(minStack.GetMin())
}
