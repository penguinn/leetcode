package index

import (
	. "github.com/penguinn/leetcode/common"
)

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	recur(root)
}

func recur(root *TreeNode) *TreeNode {
	if root.Right == nil && root.Left == nil {
		return root
	} else if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		return recur(root.Right)
	} else if root.Left == nil {
		return recur(root.Right)
	} else {
		right := root.Right
		root.Right = root.Left
		root.Left = nil
		bottom := recur(root.Right)
		bottom.Right = right
		return recur(right)
	}
}
