package index

func lastRemaining(n int, m int) int {
	position := 0
	for i:=2;i<=n;i++ {
		position = (position+m)%i
	}

	return position
}
