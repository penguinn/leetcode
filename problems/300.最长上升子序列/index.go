package index

// 当前上升序列的序列数为i, 最大值为x, 队列长度为length，求出i和length一定时，x的最小值
func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}

	iMaxList := make([]int, length)
	for i, num := range nums {
		if i == 0 {
			iMaxList[i] = 1
		} else {
			maxElement := 0
			for j := i - 1; j >= 0; j-- {
				if num > nums[j] && iMaxList[j] > maxElement {
					maxElement = iMaxList[j]
				}
			}
			iMaxList[i] = maxElement+1
		}
	}

	max := 0
	for _, iMax := range iMaxList {
		if iMax > max {
			max = iMax
		}
	}

	return max
}
