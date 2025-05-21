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
	m := make(map[string]bool)

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != 0 {
				continue
			}
			key := fmt.Sprintf("%d-%d", i, j)
			if m[key] {
				continue
			}
			// mark horizontal
			for z := 0; z < len(matrix[i]); z++ {
				if matrix[i][z] == 0 {
					continue
				}
				nextK := fmt.Sprintf("%d-%d", i, z)
				matrix[i][z] = 0
				m[nextK] = true
			}
			// mark vertical
			for z := 0; z < len(matrix); z++ {
				if matrix[z][j] == 0 {
					continue
				}
				nextK := fmt.Sprintf("%d-%d", z, j)
				matrix[z][j] = 0
				m[nextK] = true
			}
		}
	}
	// printArr(matrix)
}

func Execute(matrix [][]int) {
	setZeroes(matrix)
}
