package index

// 用的快排的思想，分为两堆。如果partition的值等于k，则直接返回；如果小于则在左边找；如果大于则在右边找
//func findKthLargest(nums []int, k int) int {
//	length := len(nums)
//	if length == 1 {
//		return nums[0]
//	}
//	return findKth(nums, 0, length-1, k)
//}
//
//func findKth(nums []int, start, end, k int) int {
//	if start == end {
//		return nums[start]
//	}
//	guard := nums[end]
//	i := start
//	for j := start; j <= end-1; j++ {
//		if nums[j] > guard {
//			if i != j {
//				nums[i], nums[j] = nums[j], nums[i]
//			}
//			i++
//		}
//	}
//	nums[i], nums[end] = nums[end], nums[i]
//	if i == k-1 {
//		return nums[i]
//	} else if i > k-1 {
//		return findKth(nums, start, i-1, k)
//	} else {
//		return findKth(nums, i+1, end, k)
//	}
//}

// 利用堆的思维
// 大顶堆
func findKthLargest(nums []int, k int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	for i := length/2 - 1; i >= 0; i-- {
		heap(&nums, i, length-1)
	}
	for i := length - 1; i >= length-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heap(&nums, 0, i-1)
	}
	return nums[length-k]
}

func heap(nums *[]int, start, end int) {
	left := 2*start + 1
	right := 2*start + 2
	if left > end {
		return
	}
	tmp := left
	if right <= end && (*nums)[right] > (*nums)[left] {
		tmp = right
	}
	if (*nums)[tmp] > (*nums)[start] {
		(*nums)[tmp], (*nums)[start] = (*nums)[start], (*nums)[tmp]
		heap(nums, tmp, end)
	} else {
		return
	}
}
