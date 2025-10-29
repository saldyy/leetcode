package problem3370

import (
	"strconv"
	"strings"
)

func smallestNumber(n int) int {
	bin := strconv.FormatInt(int64(n), 2)
	if !strings.Contains(bin, "0") {
		return n
	}

	bin = strings.ReplaceAll(bin, "0", "1")

	res, _ := strconv.ParseInt(bin, 2, 64)

	return int(res)
}

func Execute(a int) int {

	return smallestNumber(a)
}
