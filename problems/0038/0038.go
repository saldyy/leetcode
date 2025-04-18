package problem38

import (
	"strconv"
)

func runLengthEncoding(s string) string {
	i := len(s) - 2
	appearance := 1
	char := s[len(s)-1]
	r := ""
	for i >= 0 {
		if s[i] != char {
			r = strconv.Itoa(appearance) + string(char) + r
			appearance = 0
			char = s[i]
		}

		appearance++
    i--
	}

	r = strconv.Itoa(appearance) + string(char) + r
	return r
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	return runLengthEncoding(countAndSay(n - 1))
}

func Execute(a int) string {
	return countAndSay(a)
}
