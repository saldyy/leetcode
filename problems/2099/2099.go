package problem2099

import (
	"fmt"
	"sort"
)

func maxSubsequence(nums []int, k int) []int {
	ogArr := make([]int, len(nums))
	copy(ogArr, nums)
	indexes := []int{}
	m := make(map[int][]int)
	for index, value := range nums {
		if m[value] == nil {
			m[value] = []int{index}
		} else {
			m[value] = append(m[value], index)
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	fulfilled := false
	for i := range k {
		if m[nums[i]] == nil {
			continue
		}
		for _, j := range m[nums[i]] {
			indexes = append(indexes, j)
			fulfilled = len(indexes) == k
			if fulfilled {
				break
			}
		}
		m[nums[i]] = nil
		if fulfilled {
			break
		}
	}
	sort.Ints(indexes)

	ans := []int{}
	for _, i := range indexes {
		ans = append(ans, ogArr[i])
	}

	return ans
}

func Execute(a []int, b int) []int {

	return maxSubsequence(a, b)
}
