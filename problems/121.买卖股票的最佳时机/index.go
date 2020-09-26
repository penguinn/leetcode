package index

import (
	"math"
)

// 1. 保留数组每个位置的最小值
// 2. 用数组当前值减去最小值，保存结果
func maxProfit(prices []int) int {
	min := math.MaxInt64
	result := 0

	for _, price := range prices {
		if price < min {
			min = price
		}
		if price-min > result {
			result = price - min
		}
	}

	return result
}
