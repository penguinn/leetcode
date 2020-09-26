package index

import (
	"math"
)

// 动态规划
// 1. 最左边的一列节点等与上面节点+当前节点
// 2. i==j的节点等与左边节点+当前节点
// 3. 其他节点等与上面或左边的最小值+当前节点
func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return triangle[0][0]
	}
	for i := 1; i <= length-1; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				triangle[i][j] = triangle[i-1][j] + triangle[i][j]
			} else if i == j {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
			} else {
				triangle[i][j] = min(triangle[i-1][j], triangle[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	min := math.MaxInt64
	for i := 0; i <= length-1; i++ {
		if triangle[length-1][i] < min {
			min = triangle[length-1][i]
		}
	}
	return min
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
