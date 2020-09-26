package index

import (
	"math"
)

// 如果coins=[]int{1,3,5}, 那么f(7)为f(6),f(4),f(2)的最小值+1
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1, amount+1)
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if i == coin {
				dp[i] = 1
				continue
			}
			var min int
			if dp[i] != 0 {
				min = dp[i]
			} else {
				min = math.MaxInt64
			}
			if dp[i-coin] != 0 && dp[i-coin]+1 < min {
				min = dp[i-coin] + 1
			}
			if min != math.MaxInt64 {
				dp[i] = min
			}
		}
	}

	if dp[amount] == 0 {
		return -1
	} else {
		return dp[amount]
	}
}
