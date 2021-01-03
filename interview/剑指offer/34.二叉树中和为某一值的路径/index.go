package index

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var paths = [][]int{}

func pathSum(root *TreeNode, sum int) [][]int {
	paths = [][]int{}
	if root == nil {
		return paths
	}
	dfs(root, sum, []int{})
	return paths
}

func dfs(root *TreeNode, sum int, path []int) {
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			path = append(path, root.Val)
			tmpPath := make([]int, len(path))
			copy(tmpPath, path)
			paths = append(paths, tmpPath)
		}
	} else if root.Left == nil {
		sum = sum - root.Val
		path = append(path, root.Val)
		dfs(root.Right, sum, path)
	} else if root.Right == nil {
		sum = sum - root.Val
		path = append(path, root.Val)
		dfs(root.Left, sum, path)
	} else {
		sum = sum - root.Val
		path = append(path, root.Val)
		dfs(root.Left, sum, path)
		dfs(root.Right, sum, path)
	}
}
