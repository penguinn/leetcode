package index

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	charMap := map[int32]int{}
	for _, c := range s {
		if _, ok := charMap[c]; ok {
			charMap[c]++
		} else {
			charMap[c] = 1
		}
	}

	for _, c := range s {
		if charMap[c] == 1 {
			return byte(c)
		}
	}

	return ' '
}
