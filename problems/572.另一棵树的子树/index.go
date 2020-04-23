package index

import (
	"github.com/penguinn/leetcode/common"
)

func isSubtree(s *common.TreeNode, t *common.TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if (s == nil && t != nil) || (s != nil && t == nil) {
		return false
	}
	var result bool
	if isSubtree(s.Left, t) {
		result = true
	}
	if isSubtree(s.Right, t) {
		result = true
	}
	if s.Val == t.Val && isSubtreeIncludeT(s.Left, t.Left) && isSubtreeIncludeT(s.Right, t.Right) {
		result = true
	}

	return result
}

func isSubtreeIncludeT(s *common.TreeNode, t *common.TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if (s == nil && t != nil) || (s != nil && t == nil) {
		return false
	}
	if s.Val == t.Val && isSubtreeIncludeT(s.Left, t.Left) && isSubtreeIncludeT(s.Right, t.Right) {
		return true
	}
	return false
}
