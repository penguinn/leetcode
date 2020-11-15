package index

// 第1步：把所有负数变为值length+1，因为负数是作为标志的
// 第2步：根据数组中的值x，把下标x-1变为负数
// 第3步：返回数组中大于0的数的下标+1
func firstMissingPositive(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 1
	}
	for i := 0; i <= length-1; i++ {
		if nums[i] <= 0 {
			nums[i] = length + 1
		}
	}

	for i := 0; i <= length-1; i++ {
		if nums[i] != 0 && abs(nums[i]) <= length && nums[abs(nums[i])-1] > 0 {
			nums[abs(nums[i])-1] = -nums[abs(nums[i])-1]
		}
	}

	for i := 0; i <= length-1; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return length + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
