package index

func findCircleNum(M [][]int) int {
	rowLength := len(M)
	if rowLength == 0 {
		return 0
	}
	columnLength := len(M[0])
	if columnLength == 0 {
		return 0
	}
	sum := 0
	for i := 0; i <= rowLength-1; i++ {
		if M[i][i] == 1 {
			sum += 1
			M[i][i] = 0
			dfs(M, i)
		}
	}
	return sum
}

func dfs(M [][]int, i int) {
	for j := 0; j <= len(M[i])-1; j++ {
		if M[i][j] == 1 {
			M[i][j] = 0
			M[j][i] = 0
			M[j][j] = 0
			dfs(M, j)
		}
	}
}
