package main

import (
	"fmt"
)

type DoubleList struct {
	Key   string
	Value int
	Prev  *DoubleList
	Next  *DoubleList
}

type AllOne struct {
	dict   map[string]*DoubleList
	head   *DoubleList
	tail   *DoubleList
	length int
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	head := &DoubleList{}
	tail := &DoubleList{Prev: head}
	head.Next = tail

	return AllOne{
		dict:   map[string]*DoubleList{},
		head:   head,
		tail:   tail,
		length: 0,
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	node, ok := this.dict[key]
	if ok {
		node.Value++
		p := node
		for p.Next != this.tail && node.Value > p.Next.Value {
			p = p.Next
		}
		if p == node {
			return
		}
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		node.Next = p.Next
		node.Prev = p
		p.Next.Prev = node
		p.Next = node
	} else {
		tmp := &DoubleList{
			Key:   key,
			Value: 1,
			Prev:  this.head,
			Next:  this.head.Next,
		}
		this.head.Next.Prev = tmp
		this.head.Next = tmp
		this.dict[key] = tmp
		this.length++
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	node, ok := this.dict[key]
	if !ok {
		return
	}
	node.Value--
	if node.Value == 0 {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		delete(this.dict, node.Key)
		this.length--
	} else {
		p := node
		for p.Prev != this.head && node.Value < p.Prev.Value {
			p = p.Prev
		}
		if p == node {
			return
		}
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		node.Next = p
		node.Prev = p.Prev
		p.Prev.Next = node
		p.Prev = node
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.length > 0 {
		return this.tail.Prev.Key
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.length > 0 {
		return this.head.Next.Key
	}
	return ""
}


func main() {
	a := Constructor()
	a.Inc("hello")
	a.Inc("hello")
	fmt.Println(a.GetMaxKey())
	fmt.Println(a.GetMinKey())
}