package index

import (
	"math"
)

// dp[天数][交易次数][持有与否]（0；没持有， 1：持有）
func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}

	dp := make([][2 + 1][2]int, length)
	// 买入的时候算一次交易
	for i := 0; i <= 2; i++ {
		dp[0][i][1] = -prices[0]
	}

	maxK := 2
	for i := 1; i <= length-1; i++ {
		for j := maxK; j >= 1; j-- {
			dp[i][j][0] = int(math.Max(float64(dp[i-1][j][0]), float64(dp[i-1][j][1]+prices[i])))
			dp[i][j][1] = int(math.Max(float64(dp[i-1][j][1]), float64(dp[i-1][j-1][0]-prices[i])))
		}
	}

	return dp[length-1][2][0]
}
