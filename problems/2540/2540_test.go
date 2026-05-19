package problem2540

import "testing"

type Input struct {
	nums1 	[]int
	nums2 	[]int
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
				nums1: []int{1, 2, 3},
				nums2: []int{2, 4},
			},
			expected: 2,
		},
		{
			name: "Case 002",
			input: Input{
				nums1: []int{1, 2, 3, 6},
				nums2: []int{2, 3, 4, 5},
			},
			expected: 2,
		},
		{
			name: "Case 003",
			input: Input{
				nums1: []int{1, 2, 3},
				nums2: []int{4, 5, 6},
			},
			expected: -1,
		},
		{
			name: "Case 004",
			input: Input{
				nums1: []int{1, 2, 3},
				nums2: []int{4, 5, 6},
			},
			expected: -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.nums1, input.nums2)
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
