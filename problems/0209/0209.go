package problem209

import "math"

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	curSum := 0
	r := math.MaxFloat64

	for right = 0; right < len(nums); right++ {
		curSum += nums[right]
		for curSum >= target {
			r = math.Min(r, float64(right-left+1))
			curSum -= nums[left]
			left++
		}
	}

	if r == math.MaxFloat64 {
		return 0
	}
	return int(r)
}

func Execute(a int, b []int) int {

	return minSubArrayLen(a, b)
}
