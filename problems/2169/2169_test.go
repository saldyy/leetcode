package problem2169

import "testing"

type Input struct {
	num1 int
	num2 int
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
			input: Input {
				num1: 2,
				num2: 3,
			},
			expected: 3,
		},
		{
			name: "Case 002",
			input: Input {
				num1: 10,
				num2: 10,
			},
			expected: 1,
		},
		{
			name: "Case 003",
			input: Input {
				num1: 10,
				num2: 1,
			},
			expected: 10,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.num1, input.num2)
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
