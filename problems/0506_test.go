package main

import (
	"testing"
)

func TestSolution506(t *testing.T) {
	score := []int{5, 4, 3, 2, 1}

	result := Solution506(score)
	expectedVal := []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}
	for i, val := range result {
		if val != expectedVal[i] {
			t.Errorf("Expected %s but got %v", expectedVal[i], result)
		}
	}

	score = []int{10, 3, 8, 9, 4}

	result = Solution506(score)
	expectedVal = []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}
	for i, val := range result {
		if val != expectedVal[i] {
			t.Errorf("Expected %s but got %v", expectedVal[i], result)
		}
	}
}
