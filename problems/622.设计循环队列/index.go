package main

import (
	"fmt"
)

type MyCircularQueue struct {
	Queue []int
	Head  int
	Tail  int
	Count int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Queue: make([]int, k),
		Head:  0,
		Tail:  0,
		Count: 0,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.Count == len(this.Queue) {
		return false
	}
	this.Queue[this.Tail] = value
	this.Count++
	this.Tail++
	if this.Tail == len(this.Queue) {
		this.Tail = 0
	}
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.Count == 0 {
		return false
	}
	this.Count--
	this.Head++
	if this.Head == len(this.Queue) {
		this.Head = 0
	}
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.Count == 0 {
		return -1
	}
	return this.Queue[this.Head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.Count == 0 {
		return -1
	}
	var index int
	if this.Tail == 0 {
		index = len(this.Queue) - 1
	} else {
		index = this.Tail - 1
	}
	return this.Queue[index]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.Count == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.Count == len(this.Queue)
}

func main() {
	circularQueue := Constructor(6)
	fmt.Println(circularQueue.EnQueue(6))
	fmt.Println(circularQueue.Rear())
	fmt.Println(circularQueue.Rear())
	fmt.Println(circularQueue.DeQueue())
	fmt.Println(circularQueue.EnQueue(5))
	fmt.Println(circularQueue.Rear())
	fmt.Println(circularQueue.DeQueue())
	fmt.Println(circularQueue.Front())
	fmt.Println(circularQueue.DeQueue())
	fmt.Println(circularQueue.DeQueue())
	fmt.Println(circularQueue.DeQueue())
}
