package index

import . "github.com/penguinn/leetcode/common"

var maxDep int

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDep = 0
	dfs(root, 0)
	return maxDep
}

func dfs(root *TreeNode, dep int) {
	if root == nil {
		if dep > maxDep {
			maxDep = dep
		}
	} else {
		dfs(root.Left, dep+1)
		dfs(root.Right, dep+1)
	}
}
