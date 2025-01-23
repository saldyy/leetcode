package problem1267

import "testing"

type Input struct {
	grid [][]int
}

type Case struct {
	name     string
	input    Input
	expected int
}

func TestExecute(t *testing.T) {
	cases := []Case{
		{
			name:     "Case 001",
			expected: 0,
			input: Input{
				grid: [][]int{{1, 0}, {0, 1}},
			},
		},
		{
			name:     "Case 002",
			expected: 3,
			input: Input{
				grid: [][]int{{1, 0}, {1, 1}},
			},
		},
		{
			name:     "Case 003",
			expected: 4,
			input: Input{
				grid: [][]int{{1, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1}},
			},
		},
		{
			name:     "Case 004",
			expected: 3,
			input: Input{
				grid: [][]int{{1, 0, 0, 1, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 1, 0}},
			},
		},
		{
			name:     "Case 005",
			expected: 12,
			input: Input{
				grid: [][]int{{0, 0, 1, 0, 1},
					{0, 1, 0, 1, 0},
					{0, 1, 1, 1, 0},
					{1, 0, 0, 1, 1},
					{0, 0, 1, 1, 0}},
			},
		},
		{
			name:     "Case 006",
			expected: 9,
			input: Input{
				grid: [][]int{{0, 0, 0, 1, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 1, 1, 0, 0},
					{0, 1, 1, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 1, 0, 0, 0, 0, 0},
					{0, 1, 0, 1, 0, 0, 1, 0, 0}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.grid)
			if result != c.expected {
				t.Errorf("\x1b[31mFAILED %s: got=%d, want=%d\x1b[0m",
					c.name, result, c.expected)
			} else {
				t.Logf("\x1b[32mPASSED %s: got=%d\x1b[0m",
					c.name, result)
			}
		})
	}
}
