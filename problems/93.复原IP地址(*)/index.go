package index

import (
	"strconv"
	"strings"
)

const (
	IPSegmentCount    = 4
	IPSegmentMinRange = 1
	IPSegmentMaxRange = 3
)

func restoreIpAddresses(s string) []string {
	length := len(s)
	if length < IPSegmentCount*IPSegmentMinRange || length > IPSegmentCount*IPSegmentMaxRange {
		return []string{}
	}

	return dfs(0, s, 0, "")
}

func dfs(segId int, s string, sIndex int, prefix string) []string {
	result := []string{}
	if segId == 3 {
		if s[sIndex:] != "0" && strings.HasPrefix(s[sIndex:], "0") {
			return result
		}
		num, _ := strconv.Atoi(s[sIndex:])
		if num <= 255 {
			prefix = prefix + s[sIndex:]
			result = append(result, prefix)
			return result
		}
	} else {
		for i := sIndex + 1; i <= len(s)-1; i++ {
			if len(s)-i < IPSegmentMinRange*(IPSegmentCount-segId-1) || len(s)-i > IPSegmentMaxRange*(IPSegmentCount-segId-1) {
				continue
			}
			if s[sIndex:i] != "0" && strings.HasPrefix(s[sIndex:i], "0") {
				break
			}
			num, _ := strconv.Atoi(s[sIndex:i])
			if num <= 255 {
				nextPrefix := prefix + s[sIndex:i] + "."
				result = append(result, dfs(segId+1, s, i, nextPrefix)...)
			} else {
				break
			}
		}
	}
	return result
}
