package main

import (
	"fmt"
)

type MyCircularDeque struct {
	data   []int
	length int
	front  int
	last   int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	myCircularDeque := MyCircularDeque{
		data:   make([]int, k),
		length: 0,
		front:  0,
		last:   k - 1,
	}
	return myCircularDeque
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.front] = value
	this.front++
	if this.front == len(this.data) {
		this.front = 0
	}
	this.length++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.last] = value
	this.last--
	if this.last == -1 {
		this.last = len(this.data) - 1
	}
	this.length++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.front--
	if this.front == -1 {
		this.front = len(this.data) - 1
	}
	this.length--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.last++
	if this.last == len(this.data) {
		this.last = 0
	}
	this.length--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	tmp := this.front - 1
	if tmp == -1 {
		tmp = len(this.data) - 1
	}
	return this.data[tmp]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	tmp := this.last + 1
	if tmp == len(this.data) {
		tmp = 0
	}
	return this.data[tmp]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.length == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.length == len(this.data)
}

func main() {
	obj := Constructor(5)
	fmt.Println(1)
	fmt.Println(obj.InsertFront(1))
	fmt.Println(obj.InsertLast(2))
	fmt.Println(obj.DeleteFront())
	fmt.Println(obj.DeleteLast())
	fmt.Println(obj.GetFront())
	fmt.Println(obj.GetRear())
	fmt.Println(obj.IsEmpty())
	fmt.Println(obj.IsFull())
}
