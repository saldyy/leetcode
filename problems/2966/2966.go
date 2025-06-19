package problem2966

import (
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	l := len(nums) / 3
	ans := make([][]int, l)
	for i := range l {
		ans[i] = make([]int, 3)
	}
	sort.Ints(nums)

	for i, val := range nums {
		r := i/3
		c := i%3
		ans[r][c] = val
	}

	for _, arr := range ans {

		if arr[2] - arr[0] > k {
			return [][]int{}
		}
	}

	return ans
}

func Execute(a []int, b int) [][]int {

	return divideArray(a, b)
}
