package problem1957

import (
	"strings"
)

func betterVersion(s string) string {
	if len(s) < 3 {
		return s
	}

	var sb strings.Builder
	sb.Grow(len(s)) // avoid reallocations

	// Write the first two characters (safe start)
	sb.WriteByte(s[0])
	sb.WriteByte(s[1])

	for i := 2; i < len(s); i++ {
		if s[i] == s[i-1] && s[i] == s[i-2] {
			continue // skip if 3 consecutive same characters
		}
		sb.WriteByte(s[i])
	}

	return sb.String()
}

func makeFancyString(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))

	index := 0

	for index < len(s) {
		count := 1
		c := s[index]
		sb.WriteByte(c)
		for count < len(s)-index && c == s[index+count] {
			count++
			if count < 3 {
				sb.WriteByte(c)
			}
		}

		index += count
	}

	return sb.String()
}

func Execute(s string) string {

	return makeFancyString(s)
}
