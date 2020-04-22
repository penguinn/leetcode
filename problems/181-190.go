package problems

import "math"

//188.Best Time to Buy and Sell Stock IV
//按照 Best Time to Buy and Sell Stock III 的思路
//但是这里还是有个问题，如果k大于price的length，那么返回对数组可以取无限多次的算法
func maxProfit4(k int, prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	if k > length {
		return sum(prices)
	}
	i := 1
	local := make([]int, k+1)
	global := make([]int, k+1)

	for i < length {
		diff := prices[i] - prices[i-1]
		for j := k; j >= 1; j-- {
			//全局最佳，可以少加一次，当前最佳必须换算到当前节点
			if diff > 0 {
				local[j] = int(math.Max(float64(global[j-1]+diff), float64(local[j]+diff)))
			} else {
				local[j] = int(math.Max(float64(global[j-1]), float64(local[j]+diff)))
			}
			global[j] = int(math.Max(float64(global[j]), float64(local[j])))
		}
		i++
	}
	return global[k]
}

func sum(prices []int) int {
	var sum int
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			sum += diff
		}
	}
	return sum
}
