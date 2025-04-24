package problem0003

import (
)

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)

	left := 0
	r := 0

	for right := 0; right < len(s); right++ {
		if m[s[right]] != 0 {
			for s[left] != s[right] {
				delete(m, s[left])
				left++
			}
		}
		m[s[right]] = 1
		if r < len(m) {
			r = len(m)
		}
	}

	return r
}

func Execute(a string) int {

	return lengthOfLongestSubstring(a)
}
