package problem1980

import (
	"fmt"
	"strconv"
)

func find(arr []int64, n int64) bool {
	for _, val := range arr {
		if val == n {
			return true
		}
	}

	return false
}

func findDifferentBinaryString(nums []string) string {
	nis := make([]int64, len(nums))

	fmt.Printf("Len nis: %d\n", len(nis))
	for index, val := range nums {
		i, _ := strconv.ParseInt(val, 2, 64)
		fmt.Printf("I: %d\n", i)
		nis[index] = i
	}

	n := int64(0)

	if !find(nis, n) {
		return fmt.Sprintf("%0*b", len(nums), n)
	}

	n += 1

	for i := 0; i < len(nis); i++ {
		if !find(nis, n) {
			return fmt.Sprintf("%0*b", len(nums), n)
		}
		n += 1
	}

	if !find(nis, n) {
		return fmt.Sprintf("%0*b", len(nums), n)
	}

	return fmt.Sprintf("%0*b", len(nums), n+1)

}

func Execute(a []string) string {

	return findDifferentBinaryString(a)
}
