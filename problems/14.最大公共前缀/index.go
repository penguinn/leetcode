package index

// 尝试不使用额外空间
func longestCommonPrefix(strs []string) string {
	strListLength := len(strs)
	if strListLength == 0 {
		return ""
	}
	// i代表第几个字符，j代表第几个字符串
	for i := 0; i < len(strs[0]); i++ {
		for j, str := range strs {
			if len(str) == i || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
