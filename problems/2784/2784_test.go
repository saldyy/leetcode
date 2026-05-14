package problem2784

import (
	"testing"
)

type Input struct {
	nums  []int
}

type Case struct {
	name     string
	input    Input
	expected bool
}

func TestExecute(t *testing.T) {
	cases := []Case{
		{
			name: "Case 001",
			input: Input{
				nums: []int{1, 2, 3, 3},
			},
			expected: true,
		},
		{
			name: "Case 002",
			input: Input{
				nums: []int{2, 1, 3},
			},
			expected: false,
		},
		{
			name: "Case 003",
			input: Input{
				nums: []int{1, 3, 3, 2},
			},
			expected: true,
		},
		{
			name: "Case 004",
			input: Input{
				nums: []int{1, 1},
			},
			expected: true,
		},
		{
			name: "Case 005",
			input: Input{
				nums: []int{3, 4, 4, 2, 1, 1},
			},
			expected: false,
		},
		{
			name: "Case 006",
			input: Input{
				nums: []int{3, 4, 4, 2, 2, 1},
			},
			expected: false,
		},
		{
			name: "Case 007",
			input: Input{
				nums: []int{9, 9},
			},
			expected: false,
		},
		{
			name: "Case 008",
			input: Input{
				nums: []int{5, 4, 4},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.nums)
			if result != c.expected {
				t.Errorf("\x1b[31mFAILED %s: got=%v, want=%v\x1b[0m",
					c.name, result, c.expected)
			} else {
				t.Logf("\x1b[32mPASSED %s: got=%v\x1b[0m",
					c.name, result)
			}
		})
	}
}
