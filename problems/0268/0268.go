package problem268


func missingNumber(nums []int) int {
	ans := len(nums)
	for index, val := range nums {
		ans ^= val
		ans ^= index
	}

  return ans
}

func Execute(a []int) int {

	return missingNumber(a)
}
