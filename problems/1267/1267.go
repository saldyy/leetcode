package problem1267

import "fmt"

func print2D(grid [][]int) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

func countServers(grid [][]int) int {
	rowCount := make([]int, len(grid))
	colCount := make([]int, len(grid[0]))
	totalCount := 0
	visited := make([][]int, len(grid))

	for r := 0; r < len(grid); r++ {
		visited[r] = make([]int, len(grid[0]))
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 && visited[r][c] == 0 {
				visited[r][c] = 1
				rowCount[r]++
				colCount[c]++
			}
		}
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 && (rowCount[r] > 1 || colCount[c] > 1) {
				totalCount++
			}
		}
	}

	return totalCount
}

func Execute(a [][]int) int {

	return countServers(a)
}
