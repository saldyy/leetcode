package problem2206

import "fmt"

func divideArray(nums []int) bool {
	m := make(map[int]int)
	for _, val := range nums {
		m[val] = m[val] + 1
	}

  fmt.Printf("Map: %v\n", m)

  for _, val := range(m) {
    if val % 2 != 0 {
      return false
    }
  }

	return true
}

func Execute(a []int) bool {

	return divideArray(a)
}
