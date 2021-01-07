package index

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	right := searchRight(nums, target)
	left := searchLeft(nums, target)
	if right == -1 || left == -1 {
		return 0
	}
	return right - left + 1
}

func searchLeft(nums []int, target int) int {
	start := 0
	length := len(nums)
	end := length - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			if mid == 0 || nums[mid-1] < target {
				return mid
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}

func searchRight(nums []int, target int) int {
	start := 0
	length := len(nums)
	end := length - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			if mid == length-1 || nums[mid+1] > target {
				return mid
			} else {
				start = mid + 1
			}
		}
	}
	return -1
}
