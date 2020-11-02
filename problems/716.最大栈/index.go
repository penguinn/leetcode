package main

import (
	"fmt"
)

type MaxStack struct {
	stack    []int
	maxStack []int
}

/** initialize your data structure here. */
func Constructor() MaxStack {
	return MaxStack{
		stack:    []int{},
		maxStack: []int{},
	}
}

func (this *MaxStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.maxStack) == 0 || this.maxStack[len(this.maxStack)-1] <= x {
		this.maxStack = append(this.maxStack, x)
	} else {
		if this.maxStack[len(this.maxStack)-1] > x {
			this.maxStack = append(this.maxStack, this.maxStack[len(this.maxStack)-1])
		} else {
			this.maxStack = append(this.maxStack, x)
		}
	}
}

func (this *MaxStack) Pop() int {
	tmp := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.maxStack = this.maxStack[:len(this.maxStack)-1]
	return tmp
}

func (this *MaxStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MaxStack) PeekMax() int {
	return this.maxStack[len(this.maxStack)-1]
}

func (this *MaxStack) PopMax() int {
	max := this.PeekMax()
	tmpStack := []int{}
	for i := len(this.stack) - 1; i >= 0; i-- {
		if this.stack[i] != max {
			tmpStack = append(tmpStack, this.stack[i])
		} else {
			this.stack = this.stack[0:i]
			this.maxStack = this.maxStack[0:i]
			break
		}
	}
	for i := len(tmpStack) - 1; i >= 0; i-- {
		this.Push(tmpStack[i])
	}

	return max
}

func main() {
	stack := Constructor()
	stack.Push(5)
	stack.Push(1)
	fmt.Println(stack.PopMax())
	fmt.Println(stack.PeekMax())
}
