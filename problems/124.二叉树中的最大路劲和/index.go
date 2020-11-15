package index

import (
	"github.com/penguinn/leetcode/common"
	"math"
)

var maxNum int

func maxPathSum(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	maxNum = math.MinInt64
	maxPathSum_(root)
	return maxNum
}

func maxPathSum_(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	left := max(maxPathSum_(root.Left), 0)
	right := max(maxPathSum_(root.Right), 0)

	current := left + root.Val + right
	if current > maxNum {
		maxNum = current
	}

	return max(left, right) + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
