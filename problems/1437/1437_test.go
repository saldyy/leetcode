package problem1437

import "testing"

type Input struct {
	nums []int
	k    int
}

type Case struct {
	name     string
	input    Input
	expected bool
}

func TestExecute(t *testing.T) {
	cases := []Case{
		{
			name:     "Case 0001",
			input:    Input{nums: []int{1, 0, 0, 0, 1, 0, 0, 1}, k: 2},
			expected: true,
		},
		{
			name:     "Case 0002",
			input:    Input{nums: []int{1, 0, 0, 1, 0, 1}, k: 2},
			expected: false,
		},
		{
			name:     "Case 0003",
			input:    Input{nums: []int{1, 1, 0, 1, 0, 1}, k: 0},
			expected: true,
		},
		{
			name:     "Case 0004",
			input:    Input{nums: []int{0, 1, 0, 0, 1, 0, 0, 1}, k: 2},
			expected: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.nums, input.k)
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
