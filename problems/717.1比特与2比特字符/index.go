package index

func isOneBitCharacter(bits []int) bool {
	length := len(bits)
	if length == 0 {
		return false
	}

	bits = bits[:length-1]
	length = length - 1

	count := 0
	for i := 0; i <= length-1; i++ {
		if bits[i] == 1 {
			count++
		} else {
			count = 0
		}
		if count == 4 {
			return false
		}
	}
	if count == 3 {
		return false
	} else {
		return true
	}
}
