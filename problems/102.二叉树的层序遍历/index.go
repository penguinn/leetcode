package index

import (
	"github.com/penguinn/leetcode/common"
)

var dist = [][]int{}

func levelOrder(root *common.TreeNode) [][]int {
	dist = [][]int{}
	if root == nil {
		return dist
	}
	tmp := []int{root.Val}
	dist = append(dist, tmp)
	deepLevelOrder(root.Left, 1)
	deepLevelOrder(root.Right, 1)
	return dist
}

func deepLevelOrder(root *common.TreeNode, depth int) {
	if root == nil {
		return
	}
	if len(dist)-1 < depth {
		dist = append(dist, []int{root.Val})
	} else {
		dist[depth] = append(dist[depth], root.Val)
	}
	deepLevelOrder(root.Left, depth+1)
	deepLevelOrder(root.Right, depth+1)
	return
}
