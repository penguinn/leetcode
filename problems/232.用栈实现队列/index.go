package main

import (
	"fmt"
)

type MyQueue struct {
	pushStack []int
	popStack  []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		pushStack: []int{},
		popStack:  []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.pushStack = append(this.pushStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.popStack) == 0 {
		pushLength := len(this.pushStack)
		if pushLength == 0 {
			return 0
		} else {
			for i := pushLength - 1; i >= 0; i-- {
				this.popStack = append(this.popStack, this.pushStack[i])
			}
			this.pushStack = []int{}
		}
	}
	element := this.popStack[len(this.popStack)-1]
	this.popStack = this.popStack[:len(this.popStack)-1]
	return element
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.popStack) == 0 {
		pushLength := len(this.pushStack)
		if pushLength == 0 {
			return 0
		} else {
			for i := pushLength - 1; i >= 0; i-- {
				this.popStack = append(this.popStack, this.pushStack[i])
			}
			this.pushStack = []int{}
		}
	}

	return this.popStack[len(this.popStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.pushStack) == 0 && len(this.popStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())
}
