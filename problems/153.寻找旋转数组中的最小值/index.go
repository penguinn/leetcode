package index

func findMin(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	start, end := 0, length-1
	if nums[start] < nums[end] {
		return nums[start]
	}
	for start < end {
		mid := (start + end) / 2
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return nums[start]
}
