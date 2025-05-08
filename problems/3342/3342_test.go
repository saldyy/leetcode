package problem3341

import "testing"

type Input struct {
	MoveTime [][]int
}

type Case struct {
	name     string
	input    Input
	expected int
}

func TestExecute(t *testing.T) {
	cases := []Case{
		{
			name: "Case 001",
			input: Input{
				MoveTime: [][]int{{0, 4}, {4, 4}},
			},
			expected: 7,
		},
		{
			name: "Case 002",
			input: Input{
				MoveTime: [][]int{{0, 0, 0}, {0, 0, 0}},
			},
			expected: 4,
		},
		{
			name: "Case 003",
			input: Input{
				MoveTime: [][]int{{0, 1}, {1, 2}},
			},
			expected: 4,
		},
		{
			name: "Case 004",
			input: Input{
				MoveTime: [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}},
			},
			expected: 6,
		},
		{
			name: "Case 005",
			input: Input{
				MoveTime: [][]int{{21,77}, {22, 11}},
			},
			expected: 25,
		},

	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.MoveTime)
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
