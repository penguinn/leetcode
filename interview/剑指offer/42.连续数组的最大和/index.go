package index

// 尝试使用动态规划
func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	var tmp int
	cMaxNum := nums[0]
	maxNum := nums[0]
	for i := 1; i <= length-1; i++ {
		if cMaxNum > 0 {
			tmp = cMaxNum + nums[i]
		} else {
			tmp = nums[i]
		}
		cMaxNum = tmp
		if tmp > maxNum {
			maxNum = tmp
		}
	}
	return maxNum
}
