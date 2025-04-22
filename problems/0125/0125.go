package problem125

import (
	"fmt"
	"strings"
)

func formatChar(b byte) (bool, byte) {
	if b >= 48 && b <= 57 {
		return true, b
	}
	if b >= 65 && b <= 90 {
		return true, b + 32
	}
	if b >= 97 && b <= 122 {
		return true, b
	}

	return false, b
}

func isPalindrome(s string) bool {
	var sb strings.Builder
	for i := range s {
		isAscii, fb := formatChar(s[i])
		if !isAscii {
			continue
		}
		sb.WriteByte(fb)
	}

	fs := sb.String()
	left, right := 0, len(fs)-1

	fmt.Printf("S: %s\n", s)
	fmt.Printf("Formatted s: %s\n", fs)

	for left <= right {
		if fs[left] != fs[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func Execute(a string) bool {

	return isPalindrome(a)
}
