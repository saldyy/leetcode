package problem1752

import (
	"fmt"
	"sort"
)

var count = 0

func checkParts(index int, nums []int, parts [][]int) bool {
	result := false

	fmt.Println(index)
	fmt.Printf("Parts: %v\n", parts)
	fmt.Printf("Nums: %v\n", nums)
	if index > len(nums) {
		return false
	}
	if index == len(nums) {
		for i := 0; i < len(parts); i++ {
			if len(parts[i]) != 0 {
				return false
			}
		}

		return true
	}
	for i := 0; i < len(parts); i++ {
		// fmt.Printf("Parts: %v\n", parts[i])
		if len(parts[i]) > 0 && parts[i][0] == nums[index] {
			tmp := parts[i]
			parts[i] = []int{}

			nextIndex := index + len(tmp)
			if nextIndex < len(nums) && (nums[nextIndex] < tmp[len(tmp) - 1] && nums[nextIndex] > tmp[0]) {
				fmt.Printf("Found here: %v\n", tmp)
				fmt.Printf("Found here: %d, index %d\n", nums[nextIndex], nextIndex)
				continue
			}
			fmt.Printf("next index: %d\n", index+len(parts[i]))
			result = checkParts(nextIndex, nums, parts)
			parts[i] = tmp
		}

	}

	return result
}

func check(nums []int) bool {
	count := 0

	var parts [][]int
	parts = append(parts, []int{})
	parts[0] = []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			count++
			parts = append(parts, []int{nums[i]})
		} else {
			parts[count] = append(parts[count], nums[i])
		}
	}
	if (len(parts) == len(nums) && len(nums) > 2) {
		return false
	}

	sort.Ints(nums)

	fmt.Println(nums)
	fmt.Println(parts)

	return checkParts(0, nums, parts)
}

func Execute(a []int) bool {

	return check(a)
}
