package index

import (
	"github.com/penguinn/leetcode/common"
)

// 这是一种硬解的方法，从下到上的递归
func lowestCommonAncestor1(root, p, q *common.TreeNode) *common.TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	leftFind, leftNode := findAncestor(root.Left, p.Val, q.Val)
	if leftNode != nil {
		return leftNode
	}
	rightFind, rightNode := findAncestor(root.Right, p.Val, q.Val)
	if rightNode != nil {
		return rightNode
	}
	if leftFind == 1 {
		if rightFind == 1 {
			return root
		} else if root.Val == q.Val || root.Val == p.Val {
			return root
		} else {
			return nil
		}
	}
	if rightFind == 1 {
		if root.Val == q.Val || root.Val == p.Val {
			return root
		} else {
			return nil
		}
	}
	return nil
}

func findAncestor(root *common.TreeNode, p, q int) (int, *common.TreeNode) {
	if root == nil {
		return 0, nil
	}
	leftFind, leftNode := findAncestor(root.Left, p, q)
	if leftNode != nil {
		return 2, leftNode
	}
	rightFind, rightNode := findAncestor(root.Right, p, q)
	if rightNode != nil {
		return 2, rightNode
	}
	if leftFind == 1 {
		if rightFind == 1 {
			return 2, root
		} else if root.Val == q || root.Val == p {
			return 2, root
		} else {
			return 1, nil
		}
	}
	if rightFind == 1 {
		if root.Val == q || root.Val == p {
			return 2, root
		} else {
			return 1, nil
		}
	}
	if root.Val == q || root.Val == p {
		return 1, nil
	}
	return 0, nil
}

func lowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	if (root.Val <= p.Val && root.Val >= q.Val) || (root.Val >= p.Val && root.Val <= q.Val) {
		return root
	} else if leftNode := lowestCommonAncestor(root.Left, p, q); leftNode != nil {
		return leftNode
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}
