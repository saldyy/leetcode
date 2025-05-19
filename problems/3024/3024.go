package problem3024

import "sort"

func triangleType(nums []int) string {
	sort.Ints(nums)

	// Triangle inequality check
	if nums[0]+nums[1] <= nums[2] {
		return "none"
	}

	// All sides equal
	if nums[0] == nums[2] {
		return "equilateral"
	}

	// Two sides equal
	if nums[0] == nums[1] || nums[1] == nums[2] {
		return "isosceles"
	}

	// All sides different
	return "scalene"
}

func Execute(a []int) string {

	return triangleType(a)
}
