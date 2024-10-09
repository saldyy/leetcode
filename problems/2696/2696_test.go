package problem2696

import (
	"testing"
)

func Test_minLength(t *testing.T) {
	input1 := "ABFCACDB"
	result1 := minLength(input1)
	expected1 := 2

	if result1 != expected1 {
		t.Errorf("minLength(%s) = %d; expect 3", input1, result1)
	}

	input2 := "AAABBB"
	result2 := minLength(input2)
	expected2 := 0

	if result2 != expected2 {
		t.Errorf("minLength(%s) = %d; expect 3", input2, result1)
	}

}

func Test_minLengthOptimized(t *testing.T) {
	input1 := "ABFCACDB"
	result1 := minLengthOptimized(input1)
	expected1 := 2

	if result1 != expected1 {
		t.Errorf("minLengthOptimized(%s) = %d; expect 3", input1, result1)
	}

	input2 := "AAABBB"
	result2 := minLengthOptimized(input2)
	expected2 := 0 

	if result2 != expected2 {
		t.Errorf("minLengthOptimized(%s) = %d; expect 0", input2, result1)
	}

}
