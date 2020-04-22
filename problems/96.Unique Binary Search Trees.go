package problems

func numTrees(n int) int {
	list := make([]int, n+1)
	list[0] = 1
	list[1] = 1

	for flag := 2; flag <= n; flag++ {
		list[flag] = 0
		for i := 1; i < flag; i++ {
			list[flag] += list[i] * list[flag-n-1]
		}
	}
	return list[n]
}
