package index

var steps = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func maxAreaOfIsland(grid [][]int) int {
	rowLength := len(grid)
	if rowLength == 0 {
		return 0
	}
	columnLength := len(grid[0])
	max := 0
	ans := 0
	for i := 0; i <= rowLength-1; i++ {
		for j := 0; j <= columnLength-1; j++ {
			ans = dfs(grid, i, j)
			if ans > max {
				max = ans
			}
		}
	}

	return max
}

func dfs(grid [][]int, curI, curJ int) int {
	if curI == -1 || curJ == -1 || curI == len(grid) || curJ == len(grid[0]) || grid[curI][curJ] == 0 {
		return 0
	}
	grid[curI][curJ] = 0
	ans := 1

	for _, step := range steps {
		ans += dfs(grid, curI+step[0], curJ+step[1])
	}
	return ans
}

func maxAreaOfIsland1(grid [][]int) int {
	rowLength := len(grid)
	if rowLength == 0 {
		return 0
	}
	columnLength := len(grid[0])
	max := 0
	stack := [][]int{}
	ans := 0
	node := []int{}
	for i := 0; i <= rowLength-1; i++ {
		for j := 0; j <= columnLength-1; j++ {
			ans = 0
			stack = [][]int{{i, j}}
			for len(stack) > 0 {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if node[0] == -1 || node[1] == -1 || node[0] == len(grid) || node[1] == len(grid[0]) || grid[node[0]][node[1]] == 0 {
					continue
				}
				grid[node[0]][node[1]] = 0
				ans += 1
				for _, step := range steps {
					stack = append(stack, []int{node[0] + step[0], node[1] + step[1]})
				}
			}
			if ans > max {
				max = ans
			}
		}
	}

	return max
}
