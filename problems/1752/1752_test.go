package problem1752

import "testing"

type Input struct {
	nums []int
}

type Case struct {
	name     string
	input    Input
	expected bool
}

func TestExecute(t *testing.T) {
	cases := []Case{
		{
			name:     "Case 001",
			expected: true,
			input: Input{
				nums: []int{3, 4, 5, 1, 2},
			},
		},
		// {
		// 	name:     "Case 002",
		// 	expected: false,
		// 	input: Input{
		// 		nums: []int{2, 1, 3, 4},
		// 	},
		// },
		// {
		// 	name:     "Case 003",
		// 	expected: true,
		// 	input: Input{
		// 		nums: []int{1, 2, 3},
		// 	},
		// },
		// {
		// 	name:     "Case 004",
		// 	expected: true,
		// 	input: Input{
		// 		nums: []int{2, 2, 2, 2, 3, 1, 2},
		// 	},
		// },
		// {
		// 	name:     "Case 005",
		// 	expected: true,
		// 	input: Input{
		// 		nums: []int{6, 10, 6},
		// 	},
		// },
		{
			name:     "Case 006",
			expected: false,
			input: Input{
				nums: []int{11, 11, 1, 20},
			},
		},
		{
			name:     "Case 007",
			expected: true,
			input: Input{
				nums: []int{1, 2, 1, 1},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.nums)
			if result != c.expected {
				t.Errorf("\x1b[31mFAILED %s: got=%v, want=%v\x1b[0m",
					c.name, result, c.expected)
			} else {
				t.Logf("\x1b[32mPASSED %v: got=%v\x1b[0m",
					c.name, result)
			}
		})
	}
}
