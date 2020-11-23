package index

func findLengthOfLCIS(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	end := 0
	current := 0
	max := 0
	for end <= length-1 {
		if end == 0 {
			end++
			current++
		} else {
			if nums[end] > nums[end-1] {
				end++
				current++
			} else {
				current = 1
				end = end + 1
			}
		}
		if current > max {
			max = current
		}
	}

	return max
}
