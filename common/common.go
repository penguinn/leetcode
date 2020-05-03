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
