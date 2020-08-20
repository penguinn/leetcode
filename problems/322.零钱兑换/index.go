package index

import (
	"math"
)

// 如果coins=[]int{1,3,5}, 那么f(7)为f(6),f(4),f(2)的最小值+1
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	resultArray := make([]int, amount+1, amount+1)
	for i := 1; i <= amount; i++ {
		minNum := math.MaxInt64
		for _, coin := range coins {
			if i == coin {
				minNum = 0
			} else if i > coin {
				tmp := resultArray[i-coin]
				if tmp != 0 && tmp < minNum {
					minNum = tmp
				}
			}
		}
		if minNum != math.MaxInt64 {
			resultArray[i] = minNum + 1
		}
	}
	if resultArray[amount] != 0 {
		return resultArray[amount]
	} else {
		return -1
	}
}
