package index

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := equalValue(root.Left, root.Val) + equalValue(root.Right, root.Val)

	left := longestUnivaluePath(root.Left)
	if left > result {
		result = left
	}
	right := longestUnivaluePath(root.Right)
	if right > result {
		result = right
	}

	return result
}

func equalValue(root *TreeNode, value int) int {
	if root == nil {
		return 0
	}
	if root.Val != value {
		return 0
	}
	return max(equalValue(root.Left, value), equalValue(root.Right, value)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
