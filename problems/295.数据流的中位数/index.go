package main

import (
	"fmt"
)

// 插入元素自下向上堆化
// 删除元素自上向下堆化
type MedianFinder struct {
	// 大顶堆
	bigHeap   []int
	bigLength int
	// 小顶堆
	smallHeap   []int
	smallLength int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.bigLength == 0 {
		this.bigHeap = append(this.bigHeap, num)
		this.bigLength++
		this.bigHeap = bigHeapifyFromDown(this.bigHeap, this.bigLength-1, 0)
	} else if num <= this.bigHeap[0] {
		this.bigHeap = append(this.bigHeap, num)
		this.bigLength++
		this.bigHeap = bigHeapifyFromDown(this.bigHeap, this.bigLength-1, 0)
	} else {
		this.smallHeap = append(this.smallHeap, num)
		this.smallLength++
		this.smallHeap = smallHeapifyFromDown(this.smallHeap, this.smallLength-1, 0)
	}

	if this.smallLength > this.bigLength {
		tmp := this.smallHeap[0]
		this.bigHeap = append(this.bigHeap, tmp)
		this.smallHeap[0] = this.smallHeap[this.smallLength-1]
		this.smallHeap = this.smallHeap[:this.smallLength-1]
		this.smallLength--
		this.bigLength++
		this.bigHeap = bigHeapifyFromDown(this.bigHeap, this.bigLength-1, 0)
		this.smallHeap = smallHeapifyFromUp(this.smallHeap, 0, this.smallLength-1)
	}
	if this.bigLength-this.smallLength > 1 {
		tmp := this.bigHeap[0]
		this.smallHeap = append(this.smallHeap, tmp)
		this.bigHeap[0] = this.bigHeap[this.bigLength-1]
		this.bigHeap = this.bigHeap[:this.bigLength-1]
		this.smallLength++
		this.bigLength--
		this.bigHeap = bigHeapifyFromUp(this.bigHeap, 0, this.bigLength-1)
		this.smallHeap = smallHeapifyFromDown(this.smallHeap, this.smallLength-1, 0)
	}

	fmt.Println(this.bigHeap, this.bigLength)
	fmt.Println(this.smallHeap, this.smallLength)
}

// 大顶堆自上向下堆化
func bigHeapifyFromUp(nums []int, start, end int) []int {
	left := 2*start + 1
	for left <= end {
		tmp := left
		right := left + 1
		if right <= end && nums[right] > nums[tmp] {
			tmp = right
		}
		if nums[tmp] > nums[start] {
			nums[tmp], nums[start] = nums[start], nums[tmp]
			start = tmp
			left = start*2 + 1
		} else {
			break
		}
	}

	return nums
}

// 大顶堆自下向上堆化
func bigHeapifyFromDown(nums []int, start, end int) []int {
	father := (start - 1) / 2
	for father >= end && father != start {
		if nums[start] > nums[father] {
			nums[father], nums[start] = nums[start], nums[father]
			start = father
			father = (start - 1) / 2
		} else {
			break
		}
	}

	return nums
}

// 小顶堆自上向下堆化
func smallHeapifyFromUp(nums []int, start, end int) []int {
	left := 2*start + 1
	for left <= end {
		tmp := left
		right := left + 1
		if right <= end && nums[right] < nums[tmp] {
			tmp = right
		}
		if nums[tmp] < nums[start] {
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
func smallHeapifyFromDown(nums []int, start, end int) []int {
	father := (start - 1) / 2
	for father >= end && father != start {
		if nums[start] < nums[father] {
			nums[father], nums[start] = nums[start], nums[father]
			start = father
			father = (start - 1) / 2
		} else {
			break
		}
	}

	return nums
}

func (this *MedianFinder) FindMedian() float64 {
	if this.bigLength == this.smallLength && this.bigLength == 0 {
		return 0.0
	} else if this.bigLength == this.smallLength {
		return float64(this.bigHeap[0]+this.smallHeap[0]) / 2.0
	} else {
		return float64(this.bigHeap[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	medianFinder := Constructor()
	medianFinder.AddNum(40)
	medianFinder.AddNum(12)
	medianFinder.AddNum(16)
	medianFinder.AddNum(14)
	medianFinder.AddNum(35)
	medianFinder.AddNum(19)
	medianFinder.AddNum(34)
	medianFinder.AddNum(35)
	medianFinder.AddNum(28)
	medianFinder.AddNum(35)
	medianFinder.AddNum(26)
	medianFinder.AddNum(6)
	medianFinder.AddNum(8)
	medianFinder.AddNum(2)
	medianFinder.AddNum(14)
	medianFinder.AddNum(25)
	medianFinder.AddNum(25)
	medianFinder.AddNum(4)
	medianFinder.AddNum(33)
	medianFinder.AddNum(18)
	medianFinder.AddNum(10)
	medianFinder.AddNum(14)
	fmt.Println(medianFinder.FindMedian())

	fmt.Println(bigHeapifyFromUp([]int{6, 19, 26, 12, 16, 14}, 0, 5))
}

//func bigHeapifyFromUp(nums []int, start, end int) []int {
//	left := 2*start + 1
//	right := 2*start + 2
//	if left > end {
//		return nums
//	}
//	tmp := left
//	if right <= end && nums[right] > nums[tmp] {
//		tmp = right
//	}
//	if nums[tmp] > nums[start] {
//		nums[tmp], nums[start] = nums[start], nums[tmp]
//		nums = bigHeapifyFromUp(nums, tmp, end)
//	}
//	return nums
//}

//func smallHeapifyFromUp(nums []int, start, end int) []int {
//	left := 2*start + 1
//	right := 2*start + 2
//	if left > end {
//		return nums
//	}
//	tmp := left
//	if right <= end && nums[right] < nums[tmp] {
//		tmp = right
//	}
//	if nums[tmp] < nums[start] {
//		nums[tmp], nums[start] = nums[start], nums[tmp]
//		smallHeapifyFromUp(nums, tmp, end)
//	}
//	return nums
//}
