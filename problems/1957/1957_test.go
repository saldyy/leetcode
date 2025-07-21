package problem1957

import "testing"

type Input struct {
	s string
}

type Case struct {
	name     string
	input    Input
	expected string
}

func TestExecute(t *testing.T) {
	cases := []Case{
		{
			name: "case 1",
			input: Input{
				s: "leeetcode",
			},
			expected: "leetcode",
		},
		{
			name: "case 2",
			input: Input{
				s: "aaabaaaa",
			},
			expected: "aabaa",
		},
		{
			name: "case 3",
			input: Input{
				s: "aab",
			},
			expected: "aab",
		},
		{
			name: "case 4",
			input: Input{
				s: "aaaaaaa",
			},
			expected: "aa",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.s)
			if result != c.expected {
				t.Errorf(
					"\x1b[31mFAILED %s: got=%s, want=%s\x1b[0m",
					c.name,
					result,
					c.expected,
				)
			} else {
				t.Logf(
					"\x1b[32mPASSED %s: got=%s\x1b[0m",
					c.name,
					result,
				)
			}
		})
	}
}
