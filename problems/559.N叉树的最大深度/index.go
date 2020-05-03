package index

import (
	"github.com/penguinn/leetcode/common"
)

func maxDepth(root *common.Node) int {
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1
	}
	var max int
	for _, child := range root.Children {
		tmp := maxDepth(child)
		if tmp > max {
			max = tmp
		}
	}
	return max + 1
}
