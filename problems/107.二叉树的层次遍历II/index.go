package index

import (
	"github.com/penguinn/leetcode/common"
)

// BFS + 翻转
func levelOrderBottom(root *common.TreeNode) [][]int {
	result := [][]int{}
	depth := 0
	if root == nil {
		return result
	}
	result = append(result, []int{root.Val})
	traverse(root.Left, depth+1, &result)
	traverse(root.Right, depth+1, &result)
	length := len(result)
	if length <= 1 {
		return result
	}
	for i := 0; i <= (length-1)/2; i++ {
		result[i], result[length-1-i] = result[length-1-i], result[i]
	}
	return result
}

func traverse(root *common.TreeNode, depth int, result *[][]int) {
	if root == nil {
		return
	} else {
		length := len(*result)
		if depth >= length {
			*result = append(*result, []int{root.Val})
		} else {
			(*result)[depth] = append((*result)[depth], root.Val)
		}
	}
	traverse(root.Left, depth+1, result)
	traverse(root.Right, depth+1, result)
}
