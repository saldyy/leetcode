package problem1513

import "testing"

type Input struct {
	s string
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
			input:    Input{s: "0110111"},
			expected: 9,
		},
		{
			name:     "Case 002",
			input:    Input{s: "111111"},
			expected: 21,
		},
		{
			name:     "Case 003",
			input:    Input{s: "101"},
			expected: 2,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.s)
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
