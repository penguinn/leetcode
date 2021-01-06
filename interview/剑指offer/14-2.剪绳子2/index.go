package index

func cuttingRope(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	result := 1
	for n > 4 {
		result = (result * 3) % (1e9 + 7)
		n = n - 3
	}
	result = (result * n) % (1e9 + 7)
	return result
}
