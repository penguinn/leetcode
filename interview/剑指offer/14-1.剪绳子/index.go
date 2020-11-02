package index

import (
	"math"
)

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	a := n / 3
	b := n % 3
	if b == 0 {
		return int(math.Pow(3, float64(a)))
	} else if b == 1 {
		return int(math.Pow(3, float64(a-1)) * 4)
	} else {
		return int(math.Pow(3, float64(a)) * 2)
	}
}
