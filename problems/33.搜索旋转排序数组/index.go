package index

func search(nums []int, target int) int {
	length := len(nums)
	start := 0
	end := length - 1
	for start <= end {
		mid := start + ((end - start) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[start] == target {
			return start
		} else if nums[start] > target && nums[mid] > target {
			if nums[start] <= nums[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else if nums[start] > target && nums[mid] < target {
			start = mid + 1
		} else if nums[start] < target && nums[mid] > target {
			end = mid - 1
		} else {
			if nums[start] <= nums[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
