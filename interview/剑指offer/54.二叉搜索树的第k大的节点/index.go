package index

import . "github.com/penguinn/leetcode/common"

var nums []int

func kthLargest(root *TreeNode, k int) int {
	nums = []int{}
	if root == nil {
		return -1
	}
	dfs(root)
	if len(nums) < k {
		return -1
	}
	return nums[len(nums)-k]
}

func dfs(root *TreeNode) {
	if root.Left != nil {
		dfs(root.Left)
	}
	nums = append(nums, root.Val)
	if root.Right != nil {
		dfs(root.Right)
	}
}
