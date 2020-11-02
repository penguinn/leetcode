package index

import (
	"github.com/penguinn/leetcode/common"
)

// 深度遍历
func lowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}
	if root == p {
		return p
	}
	if root == q {
		return q
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left == nil && right == nil {
		return nil
	} else if left != nil {
		return left
	} else {
		return right
	}
}
