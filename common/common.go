package common

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewBinaryTree(array []interface{}) *TreeNode {
	nodeMap := map[int]*TreeNode{}

	length := len(array)
	if length <= 0 {
		return nil
	}

	for i := length - 1; i >= 0; i-- {
		if _, ok := array[i].(int); !ok {
			continue
		}
		leftIndex := 2*i + 1
		rightIndex := 2*i + 2
		leftTreeNode := nodeMap[leftIndex]
		rightTreeNode := nodeMap[rightIndex]

		treeNode := &TreeNode{
			Val:   array[i].(int),
			Left:  leftTreeNode,
			Right: rightTreeNode,
		}
		nodeMap[i] = treeNode
	}
	return nodeMap[0]
}

type Node struct {
	Val      int
	Children []*Node
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(array []interface{}) *ListNode {
	length := len(array)
	if length <= 0 {
		return nil
	}

	var list *ListNode
	p := list
	for i := 0; i < length; i++ {
		if i == 0 {
			list = &ListNode{
				Val: array[i].(int),
			}
			p = list
		} else {
			list.Next = &ListNode{
				Val: array[i].(int),
			}
			list = list.Next
		}
	}
	return p
}
