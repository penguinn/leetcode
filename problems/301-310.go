package problems

import (
	"math"
)

//309. BestTime to Buy and Shell Stock with Cooldown
//用动态规划，从后向前找，先找最高价(up)，再找出最低价(down)
//上面那种想法有漏洞，比如遇到[6,1,3,2,4,7]这样的数组，错误的想法
//这道题需要维护三个数组
//buy[i] = max(rest[i-1] - price, buy[i-1])
//rest[i] = max(rest[i-1], buy[i-1], sel[i-1])
//sel[i] = max(buy[i-1]+price , sel[i-1])
//优化
//buy[i] = max(sel[i-2] - price, buy[i-1])
//sel[i] = max(buy[i-1] + price, sel[i-1])
func maxProfit(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	buys := make([]int, length)
	sels := make([]int, length)

	for i, price := range prices {
		if i == 0 {
			buys[i] = -price
		} else if i == 1 {
			sels[i] = int(math.Max(float64(buys[i-1]+price), float64(sels[i-1])))
			buys[i] = int(math.Max(float64(-price), float64(buys[i-1])))
		} else {
			buys[i] = int(math.Max(float64(sels[i-2]-price), float64(buys[i-1])))
			sels[i] = int(math.Max(float64(buys[i-1]+price), float64(sels[i-1])))
		}
	}
	return sels[length-1]
}
