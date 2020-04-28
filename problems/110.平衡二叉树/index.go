package index

import (
	"github.com/penguinn/leetcode/common"
)

func isBalanced(root *common.TreeNode) bool {
	if root == nil {
		return true
	}
	leftBalanced, leftDepth := balanced(root.Left)
	rightBalanced, rightDepth := balanced(root.Right)
	if !leftBalanced || !rightBalanced {
		return false
	} else if leftDepth >= rightDepth && leftDepth-rightDepth <= 1 {
		return true
	} else if rightDepth >= leftDepth && rightDepth-leftDepth <= 1 {
		return true
	} else {
		return false
	}
}

func balanced(root *common.TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	} else {
		leftBalanced, leftDepth := balanced(root.Left)
		rightBalanced, rightDepth := balanced(root.Right)
		if !leftBalanced || !rightBalanced {
			return false, 0
		} else if leftDepth >= rightDepth && leftDepth-rightDepth <= 1 {
			return true, leftDepth + 1
		} else if rightDepth >= leftDepth && rightDepth-leftDepth <= 1 {
			return true, rightDepth + 1
		} else {
			return false, 0
		}
	}
}
