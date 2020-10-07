package index

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees1(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	if n == 1 {
		return []*TreeNode{{Val: 1}}
	}
	dp := make([][]*TreeNode, n+1)
	dp[0] = []*TreeNode{}
	dp[1] = []*TreeNode{{Val: 1}}

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			if len(dp[j]) == 0 {
				for _, rightNode := range dp[i-j-1] {
					dp[i] = append(dp[i], &TreeNode{Val: j + 1, Right: rightNode})
				}
				continue
			}
			if len(dp[i-j-1]) == 0 {
				for _, leftNode := range dp[j] {
					dp[i] = append(dp[i], &TreeNode{Val: j + 1, Left: leftNode})
				}
				continue
			}
			for _, leftNode := range dp[j] {
				for _, rightNode := range dp[i-j-1] {
					dp[i] = append(dp[i], &TreeNode{Val: j + 1, Left: leftNode, Right: rightNode})
				}
			}
		}
	}

	return dp[n]
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	if n == 1 {
		return []*TreeNode{{Val: 1}}
	}
	return generateTree(0, n)
}

func generateTree(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	trees := []*TreeNode{}

	for i := start; i <= end; i++ {
		leftTrees := generateTree(start, i-1)
		rightTress := generateTree(i+1, end)

		for _, left := range leftTrees {
			for _, right := range rightTress {
				trees = append(trees, &TreeNode{Val: i, Left: left, Right: right})
			}
		}
	}

	return trees
}
