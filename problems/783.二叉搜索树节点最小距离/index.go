package index

import (
	"github.com/penguinn/leetcode/common"
	"math"
)

func minDiffInBST(root *common.TreeNode) int {
	if root == nil {
		return math.MaxInt64
	}
	min := minimum(root)
	leftMin := minDiffInBST(root.Left)
	rightMin := minDiffInBST(root.Right)
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
