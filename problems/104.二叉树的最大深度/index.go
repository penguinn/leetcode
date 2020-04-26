package index

import (
	"github.com/penguinn/leetcode/common"
)

func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	max := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if right > max {
		max = right
	}
	return max + 1
}
