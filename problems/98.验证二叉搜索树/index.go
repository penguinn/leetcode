package index

import (
	"github.com/penguinn/leetcode/common"
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
