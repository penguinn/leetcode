package index

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result = [][]int{}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result = [][]int{}
	recursion(root, 0)
	for i, _ := range result {
		if i%2 == 0 {
			tmp := []int{}
			for j := 0; j <= len(result[i]) -1 ; j++ {
				tmp = append(tmp, result[i][j])
			}
			result[i] = tmp
		}
	}
	return result
}

func recursion(node *TreeNode, depth int) {
	if node == nil {
		return
	}
	if len(result) <= depth {
		result = append(result, []int{node.Val})
	} else {
		result[depth] = append(result[depth], node.Val)
	}
	recursion(node.Left, depth+1)
	recursion(node.Right, depth+1)
}
