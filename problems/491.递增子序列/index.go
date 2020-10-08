package index

import (
	"math"
)

var (
	result = [][]int{}
	tmp    = []int{}
)

func findSubsequences(nums []int) [][]int {
	result = [][]int{}
	tmp = []int{}
	dfs(0, math.MinInt64, nums)
	return result
}

func dfs(curIndex, max int, nums []int) {
	if curIndex == len(nums) {
		if len(tmp) >= 2 {
			t := make([]int, len(tmp))
			copy(t, tmp)
			result = append(result, t)
		}
		return
	}

	if nums[curIndex] >= max {
		tmp = append(tmp, nums[curIndex])
		dfs(curIndex+1, nums[curIndex], nums)
		tmp = tmp[:len(tmp)-1]
	}

	if nums[curIndex] != max {
		dfs(curIndex+1, max, nums)
	}
}
