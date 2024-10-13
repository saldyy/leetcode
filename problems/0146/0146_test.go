package problem0146

import (
	"testing"
)

func TestSolution146_1(t *testing.T) {
	cache := Constructor(1)
	cache.Put(2, 1)

	result := cache.Get(2)
	if result != 1 {
		t.Errorf("Expected 1 when getting 2 but got %v", result)
	}
}

func TestSolution146_2(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	result := cache.Get(1)
	if result != 1 {
		t.Errorf("Expected 1 when getting 1 but got %v", result)
	}

	cache.Put(3, 3)

	result = cache.Get(2)
	if result != -1 {
		t.Errorf("Expected -1 when getting 2 but got %v", result)
	}

	cache.Put(4, 4)

	result = cache.Get(1)
	if result != -1 {
		t.Errorf("Expected -1 when getting 1 but got %v", result)
	}

	result = cache.Get(3)
	if result != 3 {
		t.Errorf("Expected 3 when getting 3 but got %v", result)
	}

	result = cache.Get(4)
	if result != 4 {
		t.Errorf("Expected 4 when getting 4 but got %v", result)
	}
}

func TestSolution146_3(t *testing.T) {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)

	result := cache.Get(1)
	if result != -1 {
		t.Errorf("Expected -1 when getting 1 but got %v", result)
	}

	result = cache.Get(2)
	if result != 3 {
		t.Errorf("Expected 3 when getting 2 but got %v", result)
	}
}
