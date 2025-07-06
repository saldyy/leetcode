package problem1865

import (
	"testing"
)

func Test(t *testing.T) {
	c := Constructor(
		[]int{1, 1, 2, 2, 2, 3},
		[]int{1, 4, 5, 2, 5, 4},
	)

	r := c.Count(7)
	if r != 8 {
		t.Logf("Expected 8, got %d\n", r)
	}

	c.Add(3, 2)

	r = c.Count(8)
	if r != 2 {
		t.Errorf("Expected 2, got %d\n", r)
	}

	r = c.Count(4)
	if r != 1 {
		t.Errorf("Expected 1, got %d\n", r)
	}

	c.Add(0, 1)
	c.Add(1, 1)

	r = c.Count(7)
	if r != 11 {
		t.Errorf("Expected 11, got %d\n", r)
	}

	r = c.Count(100000)
	if r != 0 {
		t.Errorf("Expected 0, got %d\n", r)
	}

	c = Constructor(
		[]int{9, 70, 14, 9, 76},
		[]int{26, 26, 58, 23, 74, 68, 68, 78, 58, 26},
	)

	ops := [][]int{
		{6, 10},
		{5, 6},
		{32},
		{3, 55},
		{9, 32},
		{9, 16},
		{1, 48},
		{1, 4},
		{0, 52},
		{8, 20},
		{9, 4},
		{88},
		{154},
	}

	res := []int{2, 2, 7}
	i := 0
	for _, vals := range ops {
		if len(vals) == 2 {
			c.Add(vals[0], vals[1])
		} else if len(vals) == 1 {
			r := c.Count(vals[0])
			if res[i] != r {
				t.Errorf("Expected %d, got %d\n", res[i], r)
			}
			i++
		}
	}
}
