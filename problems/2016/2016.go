package problem2016

func maximumDifference(nums []int) int {
	result := -1

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			dif := nums[j] - nums[i]
			if dif > 0 && dif > result {
				result = dif
			}
		}
	}

	return result
}

func Execute(a []int) int {

	return maximumDifference(a)
}
