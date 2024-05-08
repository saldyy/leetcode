package main

import (
	"testing"

	"github.com/saldyy/leetcode/libs"
)

func TestSolution2816(t *testing.T) {
	node := &libs.Node{Value: 1}
  node.Add(8).Add(9)

	result := Solution2816(node)
	expectedVal := []int{3, 7, 8}
	i := 0
	for result != nil {
		if result.Value != expectedVal[i] {
			t.Errorf("Expected %d but got %d", expectedVal[i], result.Value)
		}
		i++
		result = result.Next
	}

	node = &libs.Node{Value: 9}
	node.Add(9).Add(9)

	result = Solution2816(node)
	expectedVal = []int{1, 9, 9, 8}
	i = 0
	for result != nil {
		if result.Value != expectedVal[i] {
			t.Errorf("Expected %d but got %d", expectedVal[i], result.Value)
		}
		i++
		result = result.Next
	}
}
