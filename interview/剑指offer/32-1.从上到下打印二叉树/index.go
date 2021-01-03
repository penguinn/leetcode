package index

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		node := nodes[len(nodes)-1]
		nodes := nodes[:len(nodes)-1]
		result = append(result, node.Val)
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}

	return result
}
