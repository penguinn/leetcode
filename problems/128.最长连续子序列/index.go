package index

func longestConsecutive(nums []int) int {
	numMap := map[int]bool{}

	for _, num := range nums {
		numMap[num] = true
	}

	max := 0
	for num := range numMap {
		current := 1
		if !numMap[num-1] {
			for numMap[num+1] {
				current++
				num++
			}
			if current > max {
				max = current
			}
		}
	}

	return max
}
