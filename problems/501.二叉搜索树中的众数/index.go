package index

import (
	"github.com/penguinn/leetcode/common"
)

func findMode(root *common.TreeNode) []int {
	nodes := ldr(root)
	if len(nodes) == 0 {
		return []int{}
	}
	returnNodes := []int{}
	count := 0
	tmpCount := 0
	for i, node := range nodes {
		if i == 0 {
			returnNodes = append(returnNodes, node)
			count = 1
			tmpCount = 1
		} else {
			if nodes[i] == nodes[i-1] {
				tmpCount += 1
				if tmpCount == count {
					returnNodes = append(returnNodes, node)
				} else if tmpCount > count {
					count = tmpCount
					returnNodes = []int{node}
				}
			} else {
				tmpCount = 1
				if tmpCount == count {
					returnNodes = append(returnNodes, node)
				}
			}
		}
	}
	return returnNodes
}

func ldr(root *common.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	lnodes := ldr(root.Left)
	lnodes = append(lnodes, root.Val)
	lnodes = append(lnodes, ldr(root.Right)...)
	return lnodes
}
