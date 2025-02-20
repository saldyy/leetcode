package problem1980

import "testing"

type Input struct {
	nums []string
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
				nums: []string{"01", "10"},
			},
			expected: "00",
		},
		{
			name: "Case 002",
			input: Input{
				nums: []string{"00", "01"},
			},
			expected: "10",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.nums)
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
