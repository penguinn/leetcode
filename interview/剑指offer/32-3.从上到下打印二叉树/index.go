package index

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var array = [][]int{}

func levelOrder(root *TreeNode) [][]int {
	array = [][]int{}
	if root == nil {
		return array
	}
	bfs(root, 1)
	for i, nums := range array {
		if i%2 == 1 {
			for j:=0;j<=len(nums)/2-1;j++ {
				nums[j], nums[len(nums)-1-j] = nums[len(nums)-1-j], nums[j]
			}
		}
	}
	return array
}

func bfs(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if len(array) < depth {
		tmp := []int{root.Val}
		array = append(array, tmp)
	} else {
		array[depth-1] = append(array[depth-1], root.Val)
	}
	bfs(root.Left, depth+1)
	bfs(root.Right, depth+1)
}
