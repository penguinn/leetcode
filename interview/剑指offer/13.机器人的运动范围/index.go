package index

var visited = [][]bool{}

func movingCount(m int, n int, k int) int {
	if m == 0 || n == 0 {
		return 0
	}
	visited = make([][]bool, 0, m)
	for i := 0; i <= m-1; i++ {
		tmp := make([]bool, n)
		visited = append(visited, tmp)
	}
	// 机器人不能跳跃
	return dfs(0, 0, m, n, k)
}

func dfs(i, j, m, n, k int) int {
	if i >= m || j >= n || visited[i][j] || get(i)+get(j) > k {
		return 0
	}
	visited[i][j] = true
	return 1 + dfs(i+1, j, m, n, k) + dfs(i, j+1, m, n, k)
}

func get(num int) int {
	var result int
	for num >= 10 {
		result += num % 10
		num = num / 10
	}
	result += num
	return result
}
