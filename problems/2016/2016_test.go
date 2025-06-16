package problem2016

import "testing"

type Input struct {
	nums []int
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
				nums: []int{7, 1, 5, 4},
			},
			expected: 4,
		},
		{
			name: "Case 002",
			input: Input{
				nums: []int{9, 4, 3, 2},
			},
			expected: -1,
		},
		{
			name: "Case 003",
			input: Input{
				nums: []int{1, 5, 2, 10},
			},
			expected: 9,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.nums)
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
