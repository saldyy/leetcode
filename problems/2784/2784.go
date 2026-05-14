package problem2784

import "slices"

func isGood(nums []int) bool {
	n := len(nums)
	if (n < 2) {
		return false
	}
	slices.Sort(nums)

	if (nums[n-1] != nums[n-2]) {
		return false
	}

	for i:=0; i<n-1; i++{
		if (nums[i] - i != 1) {
			return false
		}
	}

	return true
}

func Execute(a []int) bool {

  return isGood(a)
}
