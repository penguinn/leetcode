package problems

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	result1 := rob1(nums[0 : length-1])
	result2 := rob1(nums[1:length])
	if result1 > result2 {
		return result1
	} else {
		return result2
	}
}
