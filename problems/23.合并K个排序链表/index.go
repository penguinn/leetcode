package index

import (
	"errors"
	"github.com/penguinn/leetcode/common"
)

// 使用优先堆顶（小顶堆）
func mergeKLists(lists []*common.ListNode) *common.ListNode {
	k := len(lists)
	if k == 0 {
		return nil
	} else if k == 1 {
		return lists[0]
	} else {
		heap := NewMinHeap()
		var result *common.ListNode
		node := result
		for _, list := range lists {
			heap.Push(list)
		}
		for heap.length != 0 {
			tmp, err := heap.Pop()
			if err != nil {
				break
			}
			if node == nil {
				result = tmp
				node = tmp
			} else {
				node.Next = tmp
				node = node.Next
			}
			heap.Push(tmp.Next)
		}
		return result
	}
}

// 初始化小顶堆
type MinHeap struct {
	data   []*common.ListNode
	length int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{data: []*common.ListNode{}, length: 0}
}

func (this *MinHeap) Pop() (*common.ListNode, error) {
	if this.length == 0 {
		return nil, errors.New("heap has no elements")
	} else if this.length == 1 {
		result := this.data[0]
		this.length--
		this.data = this.data[1:]
		return result, nil
	} else {
		result := this.data[0]
		this.data[0], this.data[this.length-1] = this.data[this.length-1], this.data[0]
		this.data = this.data[:this.length-1]
		this.length--
		this.data = minHeapifyFromUp(this.data, 0, this.length-1)
		return result, nil
	}
}

func (this *MinHeap) Push(num *common.ListNode) {
	if num == nil {
		return
	}
	this.data = append(this.data, num)
	this.length++
	this.data = minHeapifyFromDown(this.data, this.length-1, 0)
}

// 小顶堆自上向下堆化
func minHeapifyFromUp(nums []*common.ListNode, start, end int) []*common.ListNode {
	left := 2*start + 1
	for left <= end {
		tmp := left
		right := left + 1
		if right <= end && nums[right].Val < nums[tmp].Val {
			tmp = right
		}
		if nums[tmp].Val < nums[start].Val {
			nums[tmp], nums[start] = nums[start], nums[tmp]
			start = tmp
			left = start*2 + 1
		} else {
			break
		}
	}

	return nums
}

// 小顶堆自下向上堆化
func minHeapifyFromDown(nums []*common.ListNode, start, end int) []*common.ListNode {
	father := (start - 1) / 2
	for father >= end && father != start {
		if nums[start].Val < nums[father].Val {
			nums[father], nums[start] = nums[start], nums[father]
			start = father
			father = (start - 1) / 2
		} else {
			break
		}
	}

	return nums
}

// 使用分治思想
func mergeKLists1(lists []*common.ListNode) *common.ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	return merge(lists, 0, length-1)
}

func merge(lists []*common.ListNode, start, end int) *common.ListNode {
	if start == end {
		return lists[start]
	} else {
		mid := (start + end) / 2
		return mergeTwoList(merge(lists, start, mid), merge(lists, mid+1, end))
	}
}

func mergeTwoList(leftList, rightList *common.ListNode) *common.ListNode {
	head := &common.ListNode{}
	p := head
	for leftList != nil && rightList != nil {
		if leftList.Val < rightList.Val {
			p.Next = leftList
			p = p.Next
			leftList = leftList.Next
		} else {
			p.Next = rightList
			p = p.Next
			rightList = rightList.Next
		}
	}
	if leftList != nil {
		p.Next = leftList
	} else {
		p.Next = rightList
	}

	return head.Next
}
