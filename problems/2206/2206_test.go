package problem2206

import "testing"

type Input struct {
	arr []int
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
				arr: []int{2, 3, 3, 2, 2, 2},
			},
			expected: true,
		},
		{
			name: "Case 002",
			input: Input{
				arr: []int{1, 2, 3, 4},
			},
			expected: false,
		},
		{
			name: "Case 003",
			input: Input{
				arr: []int{1, 1, 1, 1},
			},
			expected: true,
		},
		{
			name: "Case 004",
			input: Input{
				arr: []int{1, 2, 1, 1},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.arr)
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
