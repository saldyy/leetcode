package problem2924

import "testing"

type Input struct {
	n     int
	edges [][]int
}

type Case struct {
	name     string
	input    Input
	expected int
}

func TestExecute(t *testing.T) {
	cases := []Case{{
		name: "Case 1",
		input: Input{
			n:     3,
			edges: [][]int{{0, 1}, {1, 2}},
		},
		expected: 0,
	}, {
		name: "Case 2",
		input: Input{
			n:     4,
			edges: [][]int{{0, 2}, {1, 3}, {1, 2}},
		},
		expected: -1,
	}, {
		name: "Case 3",
		input: Input{
			n:     5,
			edges: [][]int{{4, 1}, {4, 2}, {1, 0}, {2, 0}, {0, 3}},
		},
		expected: 4,
	}, {
		name: "Case 4",
		input: Input{
			n:     6,
			edges: [][]int{{5, 2}, {4, 1}, {4, 2}, {1, 0}, {2, 0}, {0, 3}},
		},
		expected: -1,
	}, {
		name: "Case 5",
		input: Input{
			n:     1,
			edges: [][]int{},
		},
		expected: 0,
	}, {
		name: "Case 6",
		input: Input{
			n:     3,
			edges: [][]int{{0, 1}},
		},
		expected: -1,
	}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.n, input.edges)
			if result != c.expected {
				t.Errorf("\x1b[31mFAILED %s: input=(n: %d, edges: %v), got=%d, want=%d\x1b[0m",
					c.name, input.n, input.edges, result, c.expected)
			} else {
				t.Logf("\x1b[32mPASSED %s: input=(n: %d, edges: %v), got=%d\x1b[0m",
					c.name, input.n, input.edges, result)
			}
		})
	}
}
