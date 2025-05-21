package problem73

import "fmt"

func printArr(arr [][]int) {
	fmt.Printf("---------------")
	for i := range arr {
		for j := range arr[i] {
			fmt.Printf("%3d ", arr[i][j])
		}
		fmt.Printf("\n")
	}
}

func setZeroes(matrix [][]int) {
	row := make(map[int]bool)
	col := make(map[int]bool)

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i := range row {
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}

	for j := range col {
		for i := range matrix {
			matrix[i][j] = 0
		}
	}
}

func Execute(matrix [][]int) {
	setZeroes(matrix)
}
