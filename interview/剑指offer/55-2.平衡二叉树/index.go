package index

import . "github.com/penguinn/leetcode/common"

// 只遍历一次，时间复杂度较低
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, flag := dfs(root)
	return flag
}

func dfs(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	left, leftIsBalance := dfs(root.Left)
	if !leftIsBalance {
		return 0, leftIsBalance
	}
	right, rightIsBalance := dfs(root.Right)
	if !rightIsBalance {
		return 0, rightIsBalance
	}
	if left == right {
		return left + 1, true
	} else if left > right {
		if left-right > 1 {
			return 0, false
		} else {
			return left + 1, true
		}
	} else {
		if right-left > 1 {
			return 0, false
		} else {
			return right + 1, true
		}
	}
}

// 算根节点，再递归算左和右。每个节点不只遍历一遍
