package index

import (
	"github.com/penguinn/leetcode/common"
)

func minDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 {
		return right + 1
	}
	if right == 0 {
		return left + 1
	}
	if left < right {
		return left + 1
	} else {
		return right + 1
	}
}
