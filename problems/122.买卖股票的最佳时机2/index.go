package index

// 1. 峰谷法
// 2. 大数减小数，因为我总会在最低点买入，在最高点卖出
func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	result := 0
	for i := 1; i <= length-1; i++ {
		if prices[i] > prices[i-1] {
			result += prices[i]-prices[i-1]
		}
	}

	return result
}
