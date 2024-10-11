package problem962

import "fmt"

func maxWidthRamp(nums []int) int {
	n := len(nums)
	maximumRight := make([]int, n)

	maximumRight[n-1] = nums[n-1]
	for i := n-2; i >= 0; i-- {
		maximumRight[i] = max(maximumRight[i+1], nums[i])	
	}

	result := 0
	left := 0
	right := 1
	for ;right < n && left < right; {
		if maximumRight[right] >= nums[left] {
			result = max(result, right - left)
			right++
		} else {
			left++
			right = left + 1
		}	
	}

	fmt.Printf("Nums %v\n", nums)
	fmt.Printf("Maximum Right: %v\n", maximumRight)
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("------\n")
	return result
}
