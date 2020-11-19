package index

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length <= 0 {
		return 0
	}
	var max int
	charMap := map[uint8]bool{}
	left := 0
	for right := 0; right <= length-1; right++ {
		for {
			if _, ok := charMap[s[right]]; ok {
				delete(charMap, s[left])
				left++
			} else {
				break
			}
		}
		charMap[s[right]] = true
		current := len(charMap)
		if current > max {
			max = current
		}
	}
	return max
}
