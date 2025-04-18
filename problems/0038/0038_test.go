package problem38

import "testing"

type Input struct {
  n int
}

type Case struct {
	name     string
	input    Input
	expected string
}

func TestExecute(t *testing.T) {
	cases := []Case{
    {
      name: "Case 001",
      input: Input{
        n: 1,
      },
      expected: "1",
    },
    {
      name: "Case 002",
      input: Input{
        n: 4,
      },
      expected: "1211",
    },
    {
      name: "Case 003",
      input: Input{
        n: 5,
      },
      expected: "111221",
    },
    {
      name: "Case 004",
      input: Input{
        n: 6,
      },
      expected: "312211",
    },
  }

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.n)
			if result != c.expected {
				t.Errorf("\x1b[31mFAILED %s: got=%s, want=%s\x1b[0m",
					c.name, result, c.expected)
			} else {
				t.Logf("\x1b[32mPASSED %s: got=%s\x1b[0m",
					c.name, result)
			}
		})
	}
}
