package problem2099

import (
	"reflect"
	"testing"
)

type Input struct {
	nums []int
	k    int
}

type Case struct {
	name     string
	input    Input
	expected []int
}

func TestExecute(t *testing.T) {
	cases := []Case{
		// {
		// 	name: "case 001",
		// 	input: Input{
		// 		nums: []int{2, 1, 3, 3},
		// 		k:    2,
		// 	},
		// 	expected: []int{3, 3},
		// },
		// {
		// 	name: "case 002",
		// 	input: Input{
		// 		nums: []int{-1, -2, 3, 4},
		// 		k:    3,
		// 	},
		// 	expected: []int{-1, 3, 4},
		// },
		// {
		// 	name: "case 003",
		// 	input: Input{
		// 		nums: []int{3, 4, 3, 3},
		// 		k:    2,
		// 	},
		// 	expected: []int{3, 4},
		// },
		// {
		// 	name: "case 004",
		// 	input: Input{
		// 		nums: []int{18, 3, 19, -8, 30, 22, -35, 11, 16, 18, -21, 32, -7, -6, 38, 25, -21, -1, 26, -8, -37, -39, -34, 6, -36, -3, 26, -32, 22, -20, 35, -35, -30, -8, 11, 7, -23, -9, -22, 1, 33, -6, 12, 2, 27, -27, 28, -12, 21, 12, 16, 21, 33},
		// 		k:    50,
		// 	},
		// 	expected: []int{18, 3, 19, -8, 30, 22, -35, 11, 16, 18, -21, 32, -7, -6, 38, 25, -21, -1, 26, -8, -34, 6, -3, 26, -32, 22, -20, 35, -35, -30, -8, 11, 7, -23, -9, -22, 1, 33, -6, 12, 2, 27, -27, 28, -12, 21, 12, 16, 21, 33},
		// },
		{
			name: "case 005",
			input: Input{
				nums: []int{8, 2, 8, 3},
				k:    3,
			},
			expected: []int{8, 8, 3},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.nums, input.k)
			if !reflect.DeepEqual(result, c.expected) {
				t.Errorf("\x1b[31mFAILED %s: got=%d, want=%d\x1b[0m",
					c.name, result, c.expected)
			} else {
				t.Logf("\x1b[32mPASSED %s: got=%d\x1b[0m",
					c.name, result)
			}
		})
	}
}
