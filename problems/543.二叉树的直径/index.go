package index

import (
	"github.com/penguinn/leetcode/common"
)

// 遍历二叉树，结果等于左右子树之和
func diameterOfBinaryTree(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	max := diameter(root.Left) + diameter(root.Right)
	left := diameterOfBinaryTree(root.Left)
	if left > max {
		max = left
	}
	right := diameterOfBinaryTree(root.Right)
	if right > max {
		max = right
	}
	return max
}

func diameter(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	left := diameter(root.Left)
	right := diameter(root.Right)
	if right > left {
		left = right
	}
	return left + 1
}
