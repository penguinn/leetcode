package problems

import (
	"math"
	"strconv"
)

//62. Unique Paths
//1，这题是很明显的动态规划，利用递归可以完成
//2，和我猜想的一样，可能会超出时间限制，用map来储存一些结果吧 (O(n*m))
//3，网上的答案是建立一个二维数组，相加最后得到结果，其实感觉结果是一样的（O(n*m)）
//																	 min(m-1,n-1)
//4，最经典的应该是纯数学解法，其实机器人总共走了n+m-2步，其中m-1个向下，n-1个向右 C m+n-2        (O(min(n,m)
//class Solution {
//public:
//    int uniquePaths(int m, int n) {
//        double num = 1, denom = 1;
//        int small = m > n ? n : m;
//        for (int i = 1; i <= small - 1; ++i) {
//            num *= m + n - 1 - i;
//            denom *= i;
//        }
//        return (int)(num / denom);
//    }
//};
var tmp = map[string]int{}

func uniquePaths(m int, n int) int {
	tmp = map[string]int{}
	return recursive(m, n)
}

func recursive(m int, n int) int {
	var num int
	if result, ok := tmp[strconv.Itoa(m)+"."+strconv.Itoa(n)]; ok {
		return result
	}
	if m == 1 && n == 1 {
		num = 1
	} else if m == 1 {
		num = recursive(m, n-1)
	} else if n == 1 {
		num = recursive(m-1, n)
	} else {
		num = recursive(m-1, n) + recursive(m, n-1)
	}
	tmp[strconv.Itoa(m)+"."+strconv.Itoa(n)] = num
	return num
}

//63.Unique Paths II
//1，还是用动态规划加判断
//func uniquePathsWithObstacles(obstacleGrid [][]int) int {
//	tmp = map[string]int{}
//	m := len(obstacleGrid)
//	var n int
//	if m != 0 {
//		n = len(obstacleGrid[0])
//	}
//	return uniqueRecursive(&obstacleGrid, m-1, n-1)
//}
//
//func uniqueRecursive(obstacleGrid *[][]int, m int, n int) int {
//	var num int
//	if result, ok := tmp[strconv.Itoa(m)+"."+strconv.Itoa(n)]; ok {
//		return result
//	}
//	if (*obstacleGrid)[m][n] == 1 {
//		num = 0
//	} else if m == 0 && n == 0 {
//		num = 1
//	} else if m == 0 {
//		num = uniqueRecursive(obstacleGrid, m, n-1)
//	} else if n == 0 {
//		num = uniqueRecursive(obstacleGrid, m-1, n)
//	} else {
//		num = uniqueRecursive(obstacleGrid, m-1, n) + uniqueRecursive(obstacleGrid, m, n-1)
//	}
//	tmp[strconv.Itoa(m)+"."+strconv.Itoa(n)] = num
//	return num
//}
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				result[i][j] = 0
			} else if i == 0 && j == 0 {
				result[i][j] = 1
			} else if i == 0 {
				result[i][j] = result[i][j-1]
			} else if j == 0 {
				result[i][j] = result[i-1][j]
			} else {
				result[i][j] = result[i-1][j] + result[i][j-1]
			}
		}
	}
	return result[m-1][n-1]
}

//64 Minimum Path Sum
// 1，维护一个最小路径的dep[][]二维表
// 2, 生成dep[0][i]和dep[i][0]的数
// 3, 生成dep[i][j]的数
func minPathSum(grid [][]int) int {
	x := len(grid)
	if x == 0 {
		return 0
	}
	y := len(grid[0])
	if y == 0 {
		return 0
	}

	var dep [][]int

	for i := 0; i < x; i++ {
		tmp := make([]int, y)
		dep = append(dep, tmp)
	}

	dep[0][0] = grid[0][0]

	for i := 1; i < x; i++ {
		dep[i][0] = dep[i-1][0] + grid[i][0]
	}
	for j := 1; j < y; j++ {
		dep[0][j] = dep[0][j-1] + grid[0][j]
	}

	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			dep[i][j] = int(math.Min(float64(dep[i-1][j]), float64(dep[i][j-1]))) + grid[i][j]
		}
	}

	return dep[x-1][y-1]
}
