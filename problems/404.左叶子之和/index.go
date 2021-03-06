package index

import (
	"github.com/penguinn/leetcode/common"
)

/**
 * Definition for a binary tree node. */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		res += root.Left.Val
	}
	return sumOfLeftLeaves(root.Left) + res + sumOfLeftLeaves(root.Right)
}
