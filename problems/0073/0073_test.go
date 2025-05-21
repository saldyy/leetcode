package problem73

import (
	"gotest.tools/v3/assert"
	"testing"
)

type Input struct {
	matrix [][]int
}

type Case struct {
	name     string
	input    Input
	expected [][]int
}

func TestExecute(t *testing.T) {
	cases := []Case{
		// {
		// 	name: "Case 001",
		// 	input: Input{
		// 		matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
		// 	},
		// 	expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		// },
		{
			name: "Case 002",
			input: Input{
				matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			},
			expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			Execute(input.matrix)
			assert.DeepEqual(t, input.matrix, c.expected)
		})
	}
}
