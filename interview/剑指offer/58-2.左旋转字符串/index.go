package index

func reverseLeftWords(s string, n int) string {
	length := len(s)
	if length < n {
		return ""
	}
	return s[n:] + s[:n]
}
