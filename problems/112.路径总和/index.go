package index

import (
	"github.com/penguinn/leetcode/common"
)

func hasPathSum(root *common.TreeNode, sum int) bool {
	if root == nil {
		return false
	} else if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return true
		} else {
			return false
		}
	} else if root.Left == nil {
		return hasPathSum(root.Right, sum-root.Val)
	} else if root.Right == nil {
		return hasPathSum(root.Left, sum-root.Val)
	} else {
		return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
	}
}

func hasPathSum1(root *common.TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}

	return hasPathSum1(root.Left, sum-root.Val) || hasPathSum1(root.Right, sum-root.Val)
}
