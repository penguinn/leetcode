package index

var result = []string{}

func permutation(s string) []string {
	length := len(s)
	if length == 0 {
		return []string{}
	}
	if length == 1 {
		return []string{s}
	}
	result = []string{}
	dfs("", s)
	return result
}

func dfs(prefix, s string) {
	length := len(s)
	if length == 1 {
		result = append(result, prefix+s)
		return
	}

	pruningMap := map[uint8]bool{}
	for i := 0; i <= length-1; i++ {
		if _, ok := pruningMap[s[i]]; ok {
			continue
		} else {
			pruningMap[s[i]] = true
		}
		dfs(prefix+s[i:i+1], s[0:i]+s[i+1:])

	}
	return
}
