package index

func findContinuousSequence(target int) [][]int {
	if target <= 2 {
		return [][]int{}
	}

	l := 1
	r := 2

	result := [][]int{}
	for l < r {
		sum := (l + r) * (r - l + 1) / 2
		if sum == target {
			result = a(l, r, result)
			r++
		} else if sum < target {
			r++
		} else {
			l++
		}
	}
	return result
}

func a(l, r int, result [][]int) [][]int {
	list := []int{}
	for i := l; i <= r; i++ {
		list = append(list, i)
	}

	result = append(result, list)
	return result
}
