package problem209

import "testing"

type Input struct {
	target int
	nums   []int
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
				target: 7,
				nums:   []int{2, 3, 1, 2, 4, 3},
			},
			expected: 2,
		},
		{
			name: "Case 002",
			input: Input{
				target: 4,
				nums:   []int{1, 4, 4},
			},
			expected: 1,
		},
		{
			name: "Case 003",
			input: Input{
				target: 11,
				nums:   []int{1, 4, 4},
			},
			expected: 0,
		},
		{
			name: "Case 004",
			input: Input{
				target: 15,
				nums:   []int{1, 2, 3, 4, 5},
			},
			expected: 5,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.target, input.nums)
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
