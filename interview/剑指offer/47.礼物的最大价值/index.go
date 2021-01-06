package index

func maxValue(grid [][]int) int {
	rowLength := len(grid)
	if rowLength == 0 {
		return 0
	}
	columnLength := len(grid[0])
	if columnLength == 0 {
		return 0
	}
	for i := 0; i <= rowLength-1; i++ {
		for j := 0; j <= columnLength-1; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				if grid[i][j-1] > grid[i-1][j] {
					grid[i][j] = grid[i][j] + grid[i][j-1]
				} else {
					grid[i][j] = grid[i][j] + grid[i-1][j]
				}
			}
		}
	}

	return grid[rowLength-1][columnLength-1]
}
