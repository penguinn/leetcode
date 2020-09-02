package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	Iterator []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	bstIterator := BSTIterator{}
	bstIterator.iterator(root)

	return bstIterator
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	tmp := this.Iterator[len(this.Iterator)-1]
	this.Iterator = this.Iterator[:len(this.Iterator)-1]
	this.iterator(tmp.Right)
	return tmp.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.Iterator) != 0
}

func (this *BSTIterator) iterator(node *TreeNode) {
	for node != nil {
		this.Iterator = append(this.Iterator, node)
		node = node.Left
	}
}

func main() {
	iterator := Constructor(NewBinaryTree([]interface{}{7, 3, 15, nil, nil, 9, 20}))
	fmt.Println(iterator.Next())    // 返回 3
	fmt.Println(iterator.Next())    // 返回 7
	fmt.Println(iterator.HasNext()) // 返回 true
	fmt.Println(iterator.Next())    // 返回 9
	fmt.Println(iterator.HasNext()) // 返回 true
	fmt.Println(iterator.Next())    // 返回 15
	fmt.Println(iterator.HasNext()) // 返回 true
	fmt.Println(iterator.Next())    // 返回 20
	fmt.Println(iterator.HasNext()) // 返回 false
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
