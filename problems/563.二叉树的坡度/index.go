package index

import (
	"github.com/penguinn/leetcode/common"
)

func findTilt1(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	left := sum(root.Left)
	right := sum(root.Right)
	grade := left - right
	if grade < 0 {
		grade = -1 * grade
	}
	return grade + findTilt1(root.Left) + findTilt1(root.Right)
}

func sum(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	return sum(root.Left) + root.Val + sum(root.Right)
}

var title int = 0

func findTilt(root *common.TreeNode) int {
	title = 0
	traverse(root)
	return title
}

func traverse(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	left := traverse(root.Left)
	right := traverse(root.Right)
	tmp := left - right
	if tmp < 0 {
		tmp = -1 * tmp
	}
	title += tmp
	return left + right + root.Val
}
