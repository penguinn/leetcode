package index

import (
	"fmt"
	"github.com/penguinn/leetcode/common"
	"strconv"
)

func binaryTreePaths(root *common.TreeNode) []string {
	if root == nil {
		return []string{}
	} else if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	} else {
		strs := binaryTreePaths(root.Left)
		strs = append(strs, binaryTreePaths(root.Right)...)
		resultStrs := []string{}
		for _, str := range strs {
			str = fmt.Sprintf("%d->%s", root.Val, str)
			resultStrs = append(resultStrs, str)
		}
		return resultStrs
	}
}
