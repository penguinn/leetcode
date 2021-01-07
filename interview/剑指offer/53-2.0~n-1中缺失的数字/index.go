package index

// 递归思想
func missingNumber1(nums []int) int {
	return reverse(nums, 0, len(nums))
}

func reverse(nums []int, start, end int) int {
	if start == end {
		return start
	}
	mid := (start + end) / 2
	if nums[mid] == mid {
		return reverse(nums, mid+1, end)
	} else {
		return reverse(nums, start, mid)
	}
}

// 循环思想
func missingNumber(nums []int) int {
	start := 0
	end := len(nums)
	for start < end {
		mid := (start + end) / 2
		if nums[mid] == mid {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

// 若不从0开始
func NewMissingNumber(nums []int) int {
	start := 0
	end := len(nums)
	for start < end {
		mid := (start + end) / 2
		if nums[mid] == mid+nums[0] { //数组的下标数应该等于第一个数加下标（nums[1] == nums[0]+1）
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start + nums[0]
}
