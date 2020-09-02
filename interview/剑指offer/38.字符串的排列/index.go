package index

func permutation(s string) []string {
	length := len(s)
	if length == 0 {
		return []string{}
	}
	if length == 1 {
		return []string{s}
	}
	return dfs(s, "")
}

func dfs(s, prefix string) []string {
	length := len(s)
	if length == 1 {
		return []string{prefix + s}
	}
	result := []string{}
	pruningMap := map[uint8]bool{}
	for i := 0; i <= length-1; i++ {
		if _, ok := pruningMap[s[i]];ok {
			continue
		} else {
			pruningMap[s[i]] = true
		}
		if i != length-1 {
			result = append(result, dfs(s[0:i]+s[i+1:], prefix+s[i:i+1])...)
		} else {
			result = append(result, dfs(s[0:i], prefix+s[i:i+1])...)
		}
	}
	return result
}
