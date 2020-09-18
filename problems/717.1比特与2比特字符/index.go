package index

func isOneBitCharacter(bits []int) bool {
	length := len(bits)
	if length == 0 {
		return false
	}

	count := 0
	for i := 0; i <= length-2; i++ {
		if bits[i] == 1 {
			count++
		} else {
			count = 0
		}
		if count == 2 {
			count = 0
		}
	}
	if count == 1 {
		return false
	} else {
		return true
	}
}
