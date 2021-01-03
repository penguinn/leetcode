package index

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	return isSubTree(A, B)
}

func isSubTree(A *TreeNode, B *TreeNode) bool {
	if A == nil {
		return false
	}
	return sameTree(A, B) || isSubTree(A.Left, B) || isSubTree(A.Right, B)
}

func sameTree(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	} else if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return sameTree(A.Left, B.Left) && sameTree(A.Right, B.Right)
}
