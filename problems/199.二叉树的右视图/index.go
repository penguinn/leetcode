package index

import (
	"github.com/penguinn/leetcode/common"
)

func rightSideView(root *common.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}

	result = traverse(root, 0, result)

	return result
}

func traverse(root *common.TreeNode, depth int, result []int) []int {
	if root == nil {
		return result
	}
	if len(result) == depth {
		result = append(result, root.Val)
	} else {
		result[depth] = root.Val
	}
	result = traverse(root.Left, depth+1, result)
	result = traverse(root.Right, depth+1, result)
	return result
}
