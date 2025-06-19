package problem2294

import "testing"

type Input struct {
	nums []int
	k    int
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
				nums: []int{3, 2, 1},
				k: 2,
			},
			expected: 1,
		},
		{
			name: "Case 002",
			input: Input{
				nums: []int{2, 2, 4, 5},
				k: 1,
			},
			expected: 2,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.nums, input.k)
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
