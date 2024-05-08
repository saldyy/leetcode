package main

import (
	"testing"

	"github.com/saldyy/leetcode/libs"
)

func TestSolution24(t *testing.T) {
	node := &libs.Node{Value: 1}
  node.Add(2).Add(3).Add(4)
	result := Solution24(node)
	expectedVal := []int{2, 1, 4, 3}
  for _, val := range expectedVal {
    if result.Value != val {
      t.Errorf("Expected %d but got %v", val, result.Value)
    }
  }

	node = &libs.Node{Value: 1}

  node.Add(2).Add(3).Add(4).Add(5)
	result = Solution24(node)
	expectedVal = []int{2, 1, 4, 3, 5}
  for _, val := range expectedVal {
    if result.Value != val {
      t.Errorf("Expected %d but got %v", val, result.Value)
    }
  }
}
