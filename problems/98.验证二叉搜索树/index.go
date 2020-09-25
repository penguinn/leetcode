package index

import (
	"github.com/penguinn/leetcode/common"
	"math"
)

var array []int

func isValidBST(root *common.TreeNode) bool {
	if root == nil {
		return true
	}
	array = []int{}
	validBST(root.Left)
	array = append(array, root.Val)
	validBST(root.Right)
	for i, _ := range array {
		if i == 0 {
			continue
		}
		if !(array[i] > array[i-1]) {
			return false
		}
	}
	return true
}

func validBST(root *common.TreeNode) {
	if root == nil {
		return
	}
	validBST(root.Left)
	array = append(array, root.Val)
	validBST(root.Right)
}

func isValidBST1(root *common.TreeNode) bool {
	if root == nil {
		return true
	}

	return validBST1(root, math.MinInt64, math.MaxInt64)
}

func validBST1(root *common.TreeNode, lower, high int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= high {
		return false
	}

	return validBST1(root.Left, lower, root.Val) && validBST1(root.Right, root.Val, high)
}
