package index

import (
	"github.com/penguinn/leetcode/common"
)

// 用递归解法很简单，这里采用迭代，使用栈存左节点
func inorderTraversal(root *common.TreeNode) []int {
	result := []int{}
	stack := []*common.TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		index := len(stack) - 1
		root = stack[index]
		result = append(result, root.Val)
		root = root.Right
		stack = stack[:index]
	}
	return result
}
