package index


// 1. 保存了一个最大值，和最小值。
// 2. 每次根据当前值，当前值*最大值，当前值*最小值 和最大值比较
func maxProduct(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	maxMulti := nums[0]
	prevMax := nums[0]
	prevMin := nums[0]
	for i := 1; i <= length-1; i++ {
		prevMax, prevMin = maxAndMin(nums[i], nums[i]*prevMax, nums[i]*prevMin)
		if prevMax > maxMulti {
			maxMulti = prevMax
		}
	}

	return maxMulti
}

func maxAndMin(x, y, z int) (int, int) {
	var result int
	var resultMin int
	if x >= y && x >= z {
		result = x
		if y <= z {
			resultMin = y
		} else {
			resultMin = z
		}
	} else if y >= x && y >= z {
		result = y
		if x <= z {
			resultMin = x
		} else {
			resultMin = z
		}
	} else {
		result = z
		if x <= y {
			resultMin = x
		} else {
			resultMin = y
		}
	}

	return result, resultMin
}
