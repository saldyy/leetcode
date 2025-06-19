package problem2294

import (
	"sort"
)

func partitionArray(nums []int, k int) int {
	count := 1
	sort.Ints(nums)


	index := 0
	m := nums[index]
	for i := 1; i < len(nums); i++ {
		if nums[i] - m > k {
			count++
			m = nums[i]
		}
	}

	return count
}

func Execute(a []int, b int) int {

	return partitionArray(a, b)
}
