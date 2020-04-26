package index

import (
	"github.com/penguinn/leetcode/common"
)

// 可以用中序遍历，也可以用递归
func isSymmetric(root *common.TreeNode) bool {
	if root == nil {
		return true
	} else {
		return isSame(root.Left, root.Right)
	}
}

func isSame(p, q *common.TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	} else {
		return p.Val == q.Val && isSame(p.Left, q.Right) && isSame(p.Right, q.Left)
	}
}
