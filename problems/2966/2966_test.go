package problem2966

import (
	"reflect"
	"testing"
)

type Input struct {
	nums []int
	k    int
}

type Case struct {
	name     string
	input    Input
	expected [][]int
}

func TestExecute(t *testing.T) {
	cases := []Case{
		{
			name: "Case 001",
			input: Input{
				nums: []int{1, 3, 4, 8, 7, 9, 3, 5, 1},
				k:    2,
			},
			expected: [][]int{{1, 1, 3}, {3, 4, 5}, {7, 8, 9}},
		},
		{
			name: "Case 002",
			input: Input{
				nums: []int{2,4,2,2,5,2},
				k:    2,
			},
			expected: [][]int{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.nums, input.k)
			if !reflect.DeepEqual(result, c.expected) {
				t.Errorf("\x1b[31mFAILED %s: got=%d, want=%d\x1b[0m",
					c.name, result, c.expected)
			} else {
				t.Logf("\x1b[32mPASSED %s: got=%d\x1b[0m",
					c.name, result)
			}
		})
	}
}
