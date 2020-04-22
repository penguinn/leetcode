package index

import (
	"github.com/penguinn/leetcode/common"
)

// 先遍历树，在获取和为sum的
func pathSum(root *common.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return pathSum(root.Left, sum) + calculate(root, sum) + pathSum(root.Right, sum)
}

func calculate(root *common.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var count int
	if root.Val == sum {
		count += 1
	}
	return calculate(root.Left, sum-root.Val) + count + calculate(root.Right, sum-root.Val)
}
