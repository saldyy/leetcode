package problem3289

import (
	"reflect"
	"sort"
	"testing"
)

type Input struct {
	nums []int
}

type Case struct {
	name     string
	input    Input
	expected []int
}

func TestExecute(t *testing.T) {
	cases := []Case{
		{
			name: "Case 1",
			input: Input{
				nums: []int{0, 1, 1, 0},
			},
			expected: []int{0, 1},
		},
		{
			name: "Case 2",
			input: Input{
				nums: []int{1, 2, 3, 4, 2, 3},
			},
			expected: []int{2, 3},
		},
		{
			name: "Case 3",
			input: Input{
				nums: []int{0, 3, 2, 1, 3, 2},
			},
			expected: []int{2, 3},
		},
		{
			name: "Case 4",
			input: Input{
				nums: []int{7,1,5,4,3,4,6,0,9,5,8,2},
			},
			expected: []int{4, 5},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.nums)
			sort.Ints(result)

			if reflect.DeepEqual(result, c.expected) == false {
				t.Errorf("\x1b[31mFAILED %s: got=%d, want=%d\x1b[0m",
					c.name, result, c.expected)
			} else {
				t.Logf("\x1b[32mPASSED %s: got=%d\x1b[0m",
					c.name, result)
			}
		})
	}
}
