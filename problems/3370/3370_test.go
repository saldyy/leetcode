package problem3370

import "testing"

type Input struct {
	n int
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
			input:    Input{n: 5},
			expected: 7,
		},
		{
			name:     "Case 002",
			input:    Input{n: 10},
			expected: 15,
		},
		{
			name:     "Case 003",
			input:    Input{n: 3},
			expected: 3,
		},
		{
			name:     "Case 003",
			input:    Input{n: 8},
			expected: 15,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.n)
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
