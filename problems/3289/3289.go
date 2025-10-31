package problem3289

func getSneakyNumbers(nums []int) []int {
	res := []int{}
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}

	return res
}

func Execute(a []int) []int {

	return getSneakyNumbers(a)
}
