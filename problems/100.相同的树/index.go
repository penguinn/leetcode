package index

import (
	"github.com/penguinn/leetcode/common"
)

func isSameTree(p *common.TreeNode, q *common.TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	} else {
		return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
}
