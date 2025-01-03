package problem2270

func waysToSplitArray(nums []int) int {
	n := len(nums)
	leftSum := 0
	rightSum := 0
	for _, val := range nums {
		rightSum += val
	}

	count := 0
	for i := 0; i < n-1; i++ {
		leftSum += nums[i]
    rightSum -= nums[i]
		if leftSum >= rightSum {
      count++
		}
	}

	return count
}

func Execute(a []int) int {

	return waysToSplitArray(a)
}
