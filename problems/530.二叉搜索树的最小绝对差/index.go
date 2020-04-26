package index

import (
	"github.com/penguinn/leetcode/common"
	"math"
)

// 最小绝对差，先遍历树，根节点只会和它左子树的最右子树或者右子树的最左子树存在最小绝对差
func getMinimumDifference(root *common.TreeNode) int {
	if root == nil {
		return math.MaxInt64
	}
	min := minimum(root)
	leftMin := getMinimumDifference(root.Left)
	rightMin := getMinimumDifference(root.Right)
	if leftMin < min {
		min = leftMin
	}
	if rightMin < min {
		min = rightMin
	}
	return min
}

func minimum(root *common.TreeNode) int {
	min := math.MaxInt64
	rNode := root.Right
	for ; rNode != nil && rNode.Left != nil; rNode = rNode.Left {
	}
	if rNode != nil {
		min = root.Val - rNode.Val
		if min < 0 {
			min = min * -1
		}
	}

	lNode := root.Left
	for ; lNode != nil && lNode.Right != nil; lNode = lNode.Right {
	}
	tmpMin := math.MaxInt64
	if lNode != nil {
		tmpMin = root.Val - lNode.Val
		if tmpMin < 0 {
			tmpMin = tmpMin * -1
		}
	}
	if tmpMin < min {
		min = tmpMin
	}
	return min
}
