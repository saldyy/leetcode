package main

import (
	"testing"
)

func TestSolution881(t *testing.T) {
	people := []int{3, 2, 2, 1}
	limit := 3

	result := Solution881(people, limit)
	expexted := 3

	if result != expexted {
		t.Errorf("Expected %d but got %d", expexted, result)
	}

	people = []int{3, 5, 3, 4}
	limit = 5

	result = Solution881(people, limit)
	expexted = 4

	if result != expexted {
		t.Errorf("Expected %d but got %d", expexted, result)
	}

	people = []int{5, 1, 4, 2}
	limit = 6

	result = Solution881(people, limit)
	expexted = 2

	if result != expexted {
		t.Errorf("Expected %d but got %d", expexted, result)
	}
}

func BenchmarkSolution(b *testing.B) {
	people := []int{3, 2, 2, 1}
	limit := 3
	for i := 0; i < b.N; i++ {
		Solution881(people, limit)
	}
}
