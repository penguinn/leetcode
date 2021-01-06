package index

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickPow(x, n)
	}
	return 1 / quickPow(x, -n)
}

func quickPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n%2 == 1 {
		return x * quickPow(x*x, n/2)
	} else {
		return quickPow(x*x, n/2)
	}
}
