package main

import (
	"fmt"
)

type MedianFinder struct {
	minStack      []int
	minStackCount int
	maxStack      []int
	maxStackCount int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		minStack:      []int{},
		minStackCount: 0,
		maxStack:      []int{},
		maxStackCount: 0,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minStackCount == 0 && this.maxStackCount == 0 {
		this.minStack = append(this.minStack, num)
		this.minStackCount++
	} else {
		if num > this.minStack[0] {
			this.minStack = append(this.minStack, num)
			this.minStackCount++
			this.minHeapUp(this.minStackCount - 1)
		} else {
			this.maxStack = append(this.maxStack, num)
			this.maxStackCount++
			this.maxHeapUp(this.maxStackCount - 1)
		}
	}

	if this.minStackCount-this.maxStackCount > 1 {
		tmp := this.minStack[0]
		this.minStack[0], this.minStack[this.minStackCount-1] = this.minStack[this.minStackCount-1], this.minStack[0]
		this.minStackCount--
		this.minStack = this.minStack[:this.minStackCount]
		this.minHeapDown(0)
		this.maxStack = append(this.maxStack, tmp)
		this.maxStackCount++
		this.maxHeapUp(this.maxStackCount - 1)
	} else if this.maxStackCount > this.minStackCount {
		tmp := this.maxStack[0]
		this.maxStack[0], this.maxStack[this.maxStackCount-1] = this.maxStack[this.maxStackCount-1], this.maxStack[0]
		this.maxStackCount--
		this.maxStack = this.maxStack[:this.maxStackCount]
		this.maxHeapDown(0)
		this.minStack = append(this.minStack, tmp)
		this.minStackCount++
		this.minHeapUp(this.minStackCount - 1)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minStackCount == 0 {
		return 0
	}
	if this.minStackCount > this.maxStackCount {
		return float64(this.minStack[0])
	} else {
		return float64(this.minStack[0]+this.maxStack[0]) / 2
	}
}

// 小顶堆降序
func (this *MedianFinder) minHeapDown(start int) {
	left := 2*start + 1
	right := 2*start + 2

	if left > this.minStackCount-1 {
		return
	}

	tmp := left
	if right <= this.minStackCount-1 && this.minStack[right] < this.minStack[left] {
		tmp = right
	}
	if this.minStack[tmp] < this.minStack[start] {
		this.minStack[tmp], this.minStack[start] = this.minStack[start], this.minStack[tmp]
		this.minHeapDown(tmp)
	}
}

func (this *MedianFinder) minHeapUp(start int) {
	father := (start - 1) / 2
	for father >= 0 && father != start {
		if this.minStack[start] < this.minStack[father] {
			this.minStack[father], this.minStack[start] = this.minStack[start], this.minStack[father]
			start = father
			father = (start - 1) / 2
		} else {
			break
		}
	}
}

// 大顶堆
func (this *MedianFinder) maxHeapDown(start int) {
	left := 2*start + 1
	right := 2*start + 2

	if left > this.maxStackCount-1 {
		return
	}

	tmp := left
	if right <= this.maxStackCount-1 && this.maxStack[right] > this.maxStack[left] {
		tmp = right
	}
	if this.maxStack[tmp] > this.maxStack[start] {
		this.maxStack[tmp], this.maxStack[start] = this.maxStack[start], this.maxStack[tmp]
		this.maxHeapDown(tmp)
	}
}

// 大顶堆自下向上堆化
func (this *MedianFinder) maxHeapUp(start int) {
	father := (start - 1) / 2
	for father >= 0 && father != start {
		if this.maxStack[start] > this.maxStack[father] {
			this.maxStack[father], this.maxStack[start] = this.maxStack[start], this.maxStack[father]
			start = father
			father = (start - 1) / 2
		} else {
			break
		}
	}
}

func main() {
	a := Constructor()
	a.AddNum(6)
	//fmt.Println(a.FindMedian())
	a.AddNum(10)
	//fmt.Println(a.FindMedian())
	a.AddNum(2)
	//fmt.Println(a.FindMedian())
	a.AddNum(6)
	//fmt.Println(a.FindMedian())
	a.AddNum(5)
	//fmt.Println(a.FindMedian())
	a.AddNum(0)
	//fmt.Println(a.FindMedian())
	a.AddNum(6)
	fmt.Println(a.FindMedian())
	a.AddNum(3)
	//fmt.Println(a.FindMedian())
	a.AddNum(1)
	//fmt.Println(a.FindMedian())
	a.AddNum(0)
	//fmt.Println(a.FindMedian())
	a.AddNum(0)
	//fmt.Println(a.FindMedian())
}
