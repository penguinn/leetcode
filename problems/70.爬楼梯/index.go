package index

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	array := []int{1, 2}
	for i := 3; i <= n; i++ {
		tmp := array[0] + array[1]
		array[0] = array[1]
		array[1] = tmp
	}
	return array[1]
}
