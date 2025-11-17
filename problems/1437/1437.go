package problem1437

func kLengthApart(nums []int, k int) bool {
	lastOneIndex := -1
	for index, val := range nums {
		if val != 1 || lastOneIndex == -1 {
			continue
		}
		if lastOneIndex != -1 {
			l := index - lastOneIndex - 1
			if l < k {
				return false
			}
		}

		lastOneIndex = index
	}

	return true
}

func Execute(a []int, b int) bool {

	return kLengthApart(a, b)
}
