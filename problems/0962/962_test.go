package problem962

import (
	"testing"
)

func Test_maxWidthRamp(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{6, 0, 8, 2, 1, 5}, 4},
		{[]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}, 7},
		{[]int{1, 0}, 0},
		{[]int{3, 1, 2}, 1},
	}

	for _, test := range tests {
		if got := maxWidthRamp(test.nums); got != test.want {
			t.Errorf("maxWidthRamp(%v) = %v; want %v\n", test.nums, got, test.want)
		}
	}
}


