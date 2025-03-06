package problem2965

import (
	"reflect"
	"testing"
)

type Input struct {
	grid [][]int
}

type Case struct {
	name     string
	input    Input
	expected []int
}

func TestExecute(t *testing.T) {
	cases := []Case{
		{
			name:  "Case 001",
			input: Input{grid: [][]int{{1, 2}, {2, 3}}},
      expected: []int{2, 4},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.grid)
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
