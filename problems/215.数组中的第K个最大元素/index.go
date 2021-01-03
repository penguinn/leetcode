package index

// 用的快排的思想，分为两堆。如果partition的值等于k，则直接返回；如果小于则在左边找；如果大于则在右边找
func findKthLargest(nums []int, k int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	return findKth(nums, 0, length-1, k)
}

func findKth(nums []int, start, end, k int) int {
	if start == end {
		return nums[start]
	}
	guard := nums[end]
	i := start
	for j := start; j <= end-1; j++ {
		if nums[j] > guard {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	if i == k-1 {
		return nums[i]
	} else if i > k-1 {
		return findKth(nums, start, i-1, k)
	} else {
		return findKth(nums, i+1, end, k)
	}
}

// 利用堆的思维
// 小顶堆，最大长度k
func findKthLargest1(nums []int, k int) int {
	length := len(nums)
	if k > length {
		return -1
	}
	for i:=k/2-1;i>=0;i-- {
		lessHeap(nums, i, k-1)
	}

	for i:=k;i<=length-1;i++ {
		if nums[i] > nums[0] {
			nums[0], nums[i] = nums[i], nums[0]
			lessHeap(nums, 0, k-1)
		}
	}
	return nums[0]
}

func lessHeap(nums []int, start, end int) {
	left := 2*start + 1
	right := 2*start + 2
	if left > end {
		return
	}
	tmp := left
	if right <= end && nums[right] < nums[left] {
		tmp = right
	}
	if nums[tmp] < nums[start] {
		nums[tmp], nums[start] = nums[start], nums[tmp]
		lessHeap(nums, tmp, end)
	}
}
