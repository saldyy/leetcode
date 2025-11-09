package problem2169

func countOperations(num1 int, num2 int) int {
	count := 0
	for num1 != 0 && num2 != 0 {
		count++
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
		if num1 == 0 || num2 == 0 {
			return count
		}
	}

	return count
}

func Execute(a int, b int) int {

	return countOperations(a, b)
}
