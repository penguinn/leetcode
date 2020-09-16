package index

func minPathSum(grid [][]int) int {
	high := len(grid)
	if high == 0 {
		return 0
	}
	long := len(grid[0])
	if long == 0 {
		return 0
	}
	current := make([]int, long)
	for i := 0;i<=long-1;i++ {
		if i == 0 {
			current[i] = grid[0][i]
		} else {
			current[i] = current[i-1] + grid[0][i]
		}
	}
	tmp := make([]int, long)
	for i := 1; i <= high-1; i++ {
		for j := 0; j <= long-1; j++ {
			if j == 0 {
				tmp[j] = current[j] + grid[i][j]
			} else {
				tmp[j] = min(tmp[j-1], current[j]) + grid[i][j]
			}
		}
		current = tmp
	}
	return current[long-1]
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
