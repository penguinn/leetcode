/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	list := []*TreeNode{}
	if n == 1 {
		return []*TreeNode{&TreeNode{Val: 1}}
	}

	for flag := 1; flag <= n; flag++ {
		result := TreeNode{Val: flag}
		leftList := recursion(1, flag-1)
		rightList := recursion(flag+1, n)
		tmpList := generateNode(&result, leftList, rightList)
		list = append(list, tmpList...)
	}

	return list
}

func recursion(start, end int) []*TreeNode {
	list := []*TreeNode{}
	if start > end {
		return []*TreeNode{}
	} else if start == end {
		return []*TreeNode{{Val: start}}
	} else {
		for i := start; i <= end; i++ {
			result := TreeNode{Val: i}
			leftList := recursion(start, i-1)
			rightList := recursion(i+1, end)
			tmpList := generateNode(&result, leftList, rightList)
			list = append(list, tmpList...)
		}
	}
	return list
}

func generateNode(node *TreeNode, leftList, rightList []*TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if len(leftList) == 0 && len(rightList) == 0 {
		list = append(list, node)
	} else if len(leftList) == 0 {
		for _, rightNode := range rightList {
			treeNode := new(TreeNode)
			treeNode.Val = node.Val
			treeNode.Right = rightNode
			list = append(list, treeNode)
		}
	} else if len(rightList) == 0 {
		for _, leftNode := range leftList {
			treeNode := new(TreeNode)
			treeNode.Val = node.Val
			treeNode.Left = leftNode
			list = append(list, treeNode)
		}
	} else {
		for _, leftNode := range leftList {
			for _, rightNode := range rightList {
				treeNode := new(TreeNode)
				treeNode.Val = node.Val
				treeNode.Left = leftNode
				treeNode.Right = rightNode
				list = append(list, treeNode)
			}
		}
	}
	return list
}
