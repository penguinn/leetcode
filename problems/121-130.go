package problems

import (
	"math"
)

//121.Best Time to Buy and Sell stock
//1, 如果i指针等于指针且（j+1）> (j), j++；如果i指针等于指针且（j+1）< (j), j++,i++; 这个思路是错的
//2, 求出一个最小值数组，然后遍历数组分别相减求出最小值（哎这个想法应该很简单的啊）
func maxProfit1(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	var max int
	min := prices[0]
	minArray := make([]int, length)
	for i, price := range prices {
		if price < min {
			min = price
		}
		minArray[i] = min
	}

	for i, price := range prices {
		tmpMax := price - minArray[i]
		if tmpMax > max {
			max = tmpMax
		}
	}
	return max
}

//122. Best Time to Buy and Shell Stock II
//1,第一个想法是只要遇到拐点，就加一此利润，这个想法是错的
//2，找到山谷和山峰，算一次利润，时间复杂度是O(N)
//3，答案还有更简单的方法，如果后一个数大于前一个数就加1
//func maxProfit2(prices []int) int {
//	length := len(prices)
//	if length < 2 {
//		return 0
//	}
//	var maxProfit int
//	i := 0
//	j := 0
//	min := prices[0]
//	for j < length {
//		if i == j {
//			j++
//		} else if prices[j] >= prices[j-1] {
//			if prices[j] < min {
//				min = prices[j]
//			}
//			j++
//		} else if prices[j] < prices[j-1] {
//			maxProfit += prices[j-1] - min
//			i = j
//			min = prices[i]
//			j++
//		} else {
//			fmt.Println("error")
//		}
//	}
//	if prices[j-1] > min {
//		maxProfit += prices[j-1] - min
//	}
//	return maxProfit
//}
func maxProfit2(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

//123.Best Time to Buy and Sell Stock III
//1.准备借鉴II的答案，只取最大的两个波峰波谷,结果用两个值来存储，方案是错误的
//2.这个又看了一下攻略，还是使用动态规划来做，一个值用两个值来表示（当前最佳和全局最佳）
//i表示天，j表示次数
//local[i][j] = max(global[i-1][j] + diff>0?diff:0, local[i-1][j] + diff)
//global[i][j] = max(local[i][j] + global[i-1][j])
func maxProfit3(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	i := 1
	local := make([]int, 3)
	global := make([]int, 3)

	for i < length {
		diff := prices[i] - prices[i-1]
		for j := 2; j >= 1; j-- {
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
	return global[2]
}
