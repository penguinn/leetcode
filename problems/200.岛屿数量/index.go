package index

var result int

func numIslands(grid [][]byte) int {
	result = 0
	rowLength := len(grid)
	if rowLength == 0 {
		return result
	}
	columnLength := len(grid[0])
	if columnLength == 0 {
		return result
	}
	for i := 0; i <= rowLength-1; i++ {
		for j := 0; j <= columnLength-1; j++ {
			if grid[i][j] == '1' {
				result++
			}
			num(grid, i, j)
		}
	}
	return result
}

func num(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) {
		return
	}
	if j < 0 || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == '1' {
		grid[i][j] = '0'
		num(grid, i, j-1)
		num(grid, i, j+1)
		num(grid, i-1, j)
		num(grid, i+1, j)
	}
}
