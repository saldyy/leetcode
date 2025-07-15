package problem268

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
			name: "case 1",
			input: Input{
				nums: []int{0, 1, 3},
			},
			expected: 2,
		},
		{
			name: "case 2",
			input: Input{
				nums: []int{0, 2, 3},
			},
			expected: 1,
		},
		{
			name: "case 3",
			input: Input{
				nums: []int{0, 1},
			},
			expected: 2,
		},
		{
			name: "case 4",
			input: Input{
				nums: []int{1, 2, 3},
			},
			expected: 0,
		},
		{
			name: "case 5",
			input: Input{
				nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			},
			expected: 8,
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
