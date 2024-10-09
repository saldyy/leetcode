package problem2696

import (
	"bytes"
)

func deleteCharAtIndex(s string, index int, l int) string {
	var buffer bytes.Buffer
	for i := 0; i < len(s); i++ {
		if i >= index && i < index+l {
			continue
		}
		buffer.WriteString(string(s[i]))
	}

	return buffer.String()
}

func minLength(s string) int {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == 'A' && s[i+1] == 'B' {
			subS := deleteCharAtIndex(s, i, 2)
			return minLength(subS)
		}
		if s[i] == 'C' && s[i+1] == 'D' {
			subS := deleteCharAtIndex(s, i, 2)
			return minLength(subS)
		}
	}
	return len(s)
}

func minLengthOptimized(s string) int {
	stack := make([]rune, 0)
	for _, c := range s {
		if len(stack) > 0 {
			if c == 'B' && stack[len(stack)-1] == 'A' {
				stack = stack[:len(stack)-1]
				continue
			}
			if c == 'D' && stack[len(stack)-1] == 'C' {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, c)
	}

	return len(stack)
}
