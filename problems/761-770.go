package problems

import (
	"math"
)

//764.Largest Plus Sign
//这道题居然没做出来，感觉自己没有认真想啊，也是一道动态规划体啊，居然去看了一下别人的方案，但是感觉答案都是暴利解法，O(n2)还是挺慢的
//求出每个点到上下左右的有1的个数
func orderOfLargestPlusSign(N int, mines [][]int) int {
	numMap := make([][]int, N)
	for i, _ := range numMap {
		line := make([]int, N)
		numMap[i] = line
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			numMap[i][j] = 1
		}
	}
	for _, mine := range mines {
		numMap[mine[0]][mine[1]] = 0
	}

	left := make([][]int, N)
	for i, _ := range numMap {
		line := make([]int, N)
		left[i] = line
	}
	right := make([][]int, N)
	for i, _ := range numMap {
		line := make([]int, N)
		right[i] = line
	}
	up := make([][]int, N)
	for i, _ := range numMap {
		line := make([]int, N)
		up[i] = line
	}
	down := make([][]int, N)
	for i, _ := range numMap {
		line := make([]int, N)
		down[i] = line
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if numMap[i][j] == 0 {
				left[i][j] = 0
				up[i][j] = 0
			} else if i == 0 && j == 0 {
				left[i][j] = 1
				up[i][j] = 1
			} else if i == 0 {
				left[i][j] = 1
				up[i][j] = up[i][j-1] + 1
			} else if j == 0 {
				left[i][j] = left[i-1][j] + 1
				up[i][j] = 1
			} else {
				left[i][j] = left[i-1][j] + 1
				up[i][j] = up[i][j-1] + 1
			}
		}
	}

	for i := N - 1; i >= 0; i-- {
		for j := N - 1; j >= 0; j-- {
			if numMap[i][j] == 0 {
				right[i][j] = 0
				down[i][j] = 0
			} else if i == N-1 && j == N-1 {
				right[i][j] = 1
				down[i][j] = 1
			} else if i == N-1 {
				right[i][j] = 1
				down[i][j] = down[i][j+1] + 1
			} else if j == N-1 {
				right[i][j] = right[i+1][j] + 1
				down[i][j] = 1
			} else {
				right[i][j] = right[i+1][j] + 1
				down[i][j] = down[i][j+1] + 1
			}
		}
	}

	var max int

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			nodeMin := math.Min(math.Min(math.Min(float64(right[i][j]), float64(left[i][j])), float64(up[i][j])), float64(down[i][j]))
			if int(nodeMin) > max {
				max = int(nodeMin)
			}
		}
	}
	return max
}
