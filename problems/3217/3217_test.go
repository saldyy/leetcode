package problem3217

import (
	"fmt"
	"testing"
)

type Input struct {
	nums []int
	head *ListNode
}

type Case struct {
	name     string
	input    Input
	expected *ListNode
}

func TestExecute(t *testing.T) {
	cases := []Case{
		{
			name: "Test Case 001",
			input: Input{
				nums: []int{2, 3},
				head: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}},
			},
			expected: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.nums, input.head)
			fmt.Println("Result:")
			for result != nil {
				fmt.Println("Node",result.Val)
				result = result.Next
			}
			t.Error("Test not implemented yet")
		})
	}
}
