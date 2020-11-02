package main

import (
	"fmt"
)

type ListNode struct {
	key  int
	val  int
	next *ListNode
	prev *ListNode
}

type LRUCache struct {
	head       *ListNode
	tail       *ListNode
	elementMap map[int]*ListNode
	capacity   int
	length     int
}

func Constructor(capacity int) LRUCache {
	head := &ListNode{}
	tail := &ListNode{prev: head}
	head.next = tail
	return LRUCache{
		head:       head,
		tail:       tail,
		elementMap: map[int]*ListNode{},
		capacity:   capacity,
		length:     0,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.elementMap[key]
	if !ok {
		return -1
	}
	node.next.prev = node.prev
	node.prev.next = node.next

	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
	node.prev = this.head

	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.elementMap[key]
	if ok {
		node.val = value
		node.next.prev = node.prev
		node.prev.next = node.next

		node.next = this.head.next
		this.head.next.prev = node
		this.head.next = node
		node.prev = this.head
	} else {
		var tmp *ListNode
		if this.length == this.capacity {
			tmp = this.tail.prev
			this.tail.prev = tmp.prev
			tmp.prev.next = this.tail
			tmp.next = nil
			tmp.prev = nil
			delete(this.elementMap, tmp.key)

			tmp = &ListNode{
				key:  key,
				val:  value,
				next: this.head.next,
				prev: this.head,
			}
			this.head.next.prev = tmp
			this.head.next = tmp
		} else {
			tmp = &ListNode{
				key:  key,
				val:  value,
				next: this.head.next,
				prev: this.head,
			}
			this.head.next.prev = tmp
			this.head.next = tmp
			this.length++
		}
		this.elementMap[key] = tmp
	}
}

func main() {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	fmt.Println(lruCache.Get(1))
	lruCache.Put(3, 3)
	fmt.Println(lruCache.Get(2))
	lruCache.Put(4, 4)
	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(3))
	fmt.Println(lruCache.Get(4))
}
