package index

import (
	"github.com/penguinn/leetcode/common"
)

func sortedArrayToBST(nums []int) *common.TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	} else if length == 1 {
		return &common.TreeNode{Val: nums[0]}
	} else {
		var mid int
		if length%2 == 1 {
			mid = (length - 1) / 2
		} else {
			mid = length / 2
		}
		return &common.TreeNode{Val: nums[mid], Left: sortedArrayToBST(nums[0:mid]), Right: sortedArrayToBST(nums[mid+1 : length])}
	}
}
