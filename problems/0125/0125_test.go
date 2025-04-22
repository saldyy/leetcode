package problem125

import "testing"

type Input struct {
	s string
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
        s: " ",
      },
      expected: true,
    },
    {
      name: "Case 002",
      input: Input{
        s: "A man, a plan, a canal: Panama",
      },
      expected: true,
    },
    {
      name: "Case 003",
      input: Input{
        s: "race a car",
      },
      expected: false,
    },
    {
      name: "Case 004",
      input: Input{
        s: "0P",
      },
      expected: false,
    },
    {
      name: "Case 005",
      input: Input{
        s: ".,",
      },
      expected: true,
    },
    {
      name: "Case 006",
      input: Input{
        s: ".,.",
      },
      expected: true,
    },
  }

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input.s)
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
