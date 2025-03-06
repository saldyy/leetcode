package problem2965

func findMissingAndRepeatedValues(grid [][]int) []int {
	result := make([]int, 2)
	m := make(map[int]int)
	r := len(grid)
	c := len(grid[0])
	s := 0

	for row := 0; row < r; row++ {
		for col := 0; col < c; col++ {
			m[grid[row][col]] = m[grid[row][col]] + 1
			s += grid[row][col]
			if m[grid[row][col]] > 1 {
				result[0] = grid[row][col]
			}
		}
	}

  
	sum := ((r*c + 1) * r * c) / 2
	result[1] = sum - (s - result[0])

	return result
}

func Execute(a [][]int) []int {

	return findMissingAndRepeatedValues(a)
}
