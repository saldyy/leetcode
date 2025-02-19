package problem1415

import (
	"fmt"
	"sort"
	"strings"
)

var ch = []string{"a", "b", "c"}

func gen(s string, n int, r *[]string) {
	l := len(s)

	if l == n {
		fmt.Printf("S: %s\n", s)
		*r = append(*r, s)
		return
	}

	for _, c := range ch {
		if s[l-1] == c[0] {
			continue
		}
		gen(s+c, n, r)
	}
}

func getHappyString(n int, k int) string {
	r := []string{}
	gen("a", n, &r)
	gen("b", n, &r)
	gen("c", n, &r)

	sort.Slice(r, func(i, j int) bool {
		return strings.ToLower(r[i]) < strings.ToLower(r[j])
	})

	fmt.Println(r)

	if k > len(r) {
		return ""
	}

	return r[k-1]
}

func Execute(a int, b int) string {

	return getHappyString(a, b)
}
