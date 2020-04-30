package index

import (
	"github.com/penguinn/leetcode/common"
)

func bstToGst(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}
	sum := traverse(root.Right, 0)
	sum = sum + root.Val
	root.Val = sum
	traverse(root.Left, sum)
	return root
}

func traverse(root *common.TreeNode, sum int) int {
	if root == nil {
		return sum
	}
	sum = traverse(root.Right, sum)
	sum = sum + root.Val
	root.Val = sum
	sum = traverse(root.Left, sum)
	return sum
}