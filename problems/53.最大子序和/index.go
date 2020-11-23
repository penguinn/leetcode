package index

import (
	"math"
)

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	max := math.MinInt64
	current := 0

	for i:=0;i<=length-1;i++ {
		if current <= 0 {
			current = nums[i]
		} else {
			current = current+nums[i]
		}
		if current > max {
			max = current
		}
	}

	return max
}
