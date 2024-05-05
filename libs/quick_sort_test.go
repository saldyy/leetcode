package libs

import (
	"testing"
)

func TestQuickSortAsc(t *testing.T) {
	arr := []int{10, 5, 2, 3}
	expected := []int{2, 3, 5, 10}
	QuickSort(arr, 0, len(arr)-1, "asc")

	for i, v := range arr {
		if v != expected[i] {
			t.Errorf("Expected %d but got %d", expected[i], v)
		}
	}
}

func TestQuickSortDesc(t *testing.T) {
	arr := []int{10, 5, 2, 3}
	expected := []int{10, 5, 3, 2}
	QuickSort(arr, 0, len(arr)-1, "desc")

	for i, v := range arr {
		if v != expected[i] {
			t.Errorf("Expected %d but got %d", expected[i], v)
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	arr := []int{10, 5, 2, 3}
	for i := 0; i < b.N; i++ {
		QuickSort(arr, 0, len(arr)-1, "asc")
	}
}
