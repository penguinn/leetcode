package problems

import (
	"strconv"
)

/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	l := len(s)
	if l == 1 {
		_, err := strconv.Atoi(s)
		if err != nil {
			return 0
		}
		return 1
	}
	ways := make([]int, l+1)
	ways[0] = 1
	if s[0] == '0' {
		ways[1] = 0
	} else {
		ways[1] = 1
	}
	for i := 2; i <= l; i++ {
		str := s[i-2 : i]
		num, err := strconv.Atoi(string(str))
		if err != nil {
			return 0
		}
		if num == 0 {
			return 0
		} else if num > 0 && num <= 9 {
			ways[i] = ways[i-1]
		} else if num == 10 || num == 20 {
			ways[i] = ways[i-2]
		} else if num == 30 || num == 40 || num == 50 || num == 60 || num == 70 || num == 80 || num == 90 {
			return 0
		} else if (num > 20 && num <= 26) || (num > 10 && num < 20) {
			ways[i] = ways[i-1] + ways[i-2]
		} else {
			ways[i] = ways[i-1]
		}
	}
	return ways[l]
}

