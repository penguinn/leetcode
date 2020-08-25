package index

func findRepeatNumber(nums []int) int {
	validMap := map[int]bool{}
	for _, num := range nums {
		if _, ok := validMap[num];ok {
			return num
		} else {
			validMap[num] = true
		}
	}

	return 0
}
