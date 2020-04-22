package problems
/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

func climbStairs(n int) int {
	ways := []int{1, 1}
	for i := 1; i < n; i++ {
		temp := ways[1]
		ways[1] += ways[0]
		ways[0] = temp
	}
	return ways[1]
}

