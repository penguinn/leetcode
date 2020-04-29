package index

import (
	"github.com/penguinn/leetcode/common"
)

func invertTree(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
