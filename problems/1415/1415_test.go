package problem1415

import "testing"

type Input struct {
	n int
	k int
}

type Case struct {
	name     string
	input    Input
	expected string
}

func TestExecute(t *testing.T) {
	cases := []Case{
		{
			name: "Case 001",
			input: Input{
				n: 1,
				k: 3,
			},
			expected: "c",
		},
		{
			name: "Case 002",
			input: Input{
				n: 3,
				k: 9,
			},
			expected: "cab",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.n, input.k)
			if result != c.expected {
				t.Errorf("\x1b[31mFAILED %s: got=%s, want=%s\x1b[0m",
					c.name, result, c.expected)
			} else {
				t.Logf("\x1b[32mPASSED %s: got=%s\x1b[0m",
					c.name, result)
			}
		})
	}
}
