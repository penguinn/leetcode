package index

import (
	"fmt"
	"github.com/penguinn/leetcode/common"
)

func buildTree(preorder []int, inorder []int) *common.TreeNode {
	len1 := len(preorder)
	len2 := len(inorder)
	if len1 != len2 {
		fmt.Println(len1, len2)
	}
	if len1 == 0 {
		return nil
	}
	if len1 == 1 {
		return &common.TreeNode{Val: preorder[0]}
	}

	val := preorder[0]
	var i int
	for ; i < len2; i++ {
		if inorder[i] == val {
			break
		}
	}

	leftLength := i - 0

	return &common.TreeNode{Val: val, Left: buildTree(preorder[1:1+leftLength], inorder[0:i]), Right: buildTree(preorder[1+leftLength:], inorder[i+1:])}
}
