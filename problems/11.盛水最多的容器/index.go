package index

func maxArea(height []int) int {
	length := len(height)
	if length <= 1 {
		return 0
	}
	max := 0
	left := 0
	right := length - 1
	var tmp int
	for left < right {
		if height[left] >= height[right] {
			tmp = height[right] * (length-1)
			right--
		} else {
			tmp = height[left] * (length-1)
			left++
		}
		length--
		if tmp > max {
			max = tmp
		}
	}

	return max
}
