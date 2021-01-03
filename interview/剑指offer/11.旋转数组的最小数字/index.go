package index

func minArray(nums []int) int {
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
		} else if nums[mid] < nums[end] {
			end = mid
		} else {
			end = end - 1
		}
	}

	return nums[start]
}
