package index

import (
	"math"
)

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
			} else {
				triangle[i][j] = min(triangle[i-1][j], triangle[i-1][j+1]) + triangle[i][j]
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
