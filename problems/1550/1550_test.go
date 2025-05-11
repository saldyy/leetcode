package problem1550

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
				arr: []int{2, 6, 4, 1},
			},
			expected: false,
		},
		{
			name: "Case 002",
			input: Input{
				arr: []int{1, 2, 34, 3, 4, 5, 7, 23, 12},
			},
			expected: true,
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
