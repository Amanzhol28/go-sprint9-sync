package main

import "testing"

func TestMaximum(t *testing.T) {
	data := []int{12, 45, 7, 89, 23, 56}

	result := maximum(data)

	if result != 89 {
		t.Errorf("Expected 89, got %d", result)
	}
}

func TestMaximumEmpty(t *testing.T) {
	data := []int{}

	result := maximum(data)

	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestMaxChunksEmpty(t *testing.T) {
	data := []int{}

	result := maxChunks(data)

	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestMaxChunks(t *testing.T) {
	data := []int{
		12, 45, 7, 89,
		23, 56, 100, 34,
		67, 91, 15, 200,
		8, 76, 33, 54,
	}

	result := maxChunks(data)

	if result != 200 {
		t.Errorf("Expected 200, got %d", result)
	}
}

func TestMaxChunksSmallSlice(t *testing.T) {
	data := []int{15, 8, 30, 4, 22}

	result := maxChunks(data)

	if result != 30 {
		t.Errorf("Expected 30, got %d", result)
	}
}

func TestGenerateRandomElements(t *testing.T) {
	data := generateRandomElements(100)

	if len(data) != 100 {
		t.Errorf("Expected length 100, got %d", len(data))
	}

	for _, num := range data {
		if num < 1 || num > 1000 {
			t.Errorf("Number out of range: %d", num)
		}
	}
}

func TestMaxChunksWithRemainder(t *testing.T) {
	data := []int{
		12, 45, 7, 89, 23,
		56, 100, 34, 67, 91,
		15, 200, 8, 76, 33,
		54, 21, 44, 999,
	}

	result := maxChunks(data)

	if result != 999 {
		t.Errorf("Expected 999, got %d", result)
	}
}
