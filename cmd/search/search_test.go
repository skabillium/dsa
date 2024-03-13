package search

import (
	"math/rand"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	haystack := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	if !LinearSearch(haystack, 69) {
		t.Error("Expected true, go false")
	}
	if !LinearSearch(haystack, 69420) {
		t.Error("Expected true, go false")
	}
	if !LinearSearch(haystack, 1) {
		t.Error("Expected true, go false")
	}

	if LinearSearch(haystack, 1336) {
		t.Error("Expected false, go true")
	}
	if LinearSearch(haystack, 69421) {
		t.Error("Expected false, go true")
	}
	if LinearSearch(haystack, 2) {
		t.Error("Expected false, go true")
	}
}

func TestBinarySearch(t *testing.T) {
	haystack := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	if !BinarySearch(haystack, 69) {
		t.Error("Expected true, go false")
	}
	if !BinarySearch(haystack, 69420) {
		t.Error("Expected true, go false")
	}
	if !BinarySearch(haystack, 1) {
		t.Error("Expected true, go false")
	}

	if BinarySearch(haystack, 1336) {
		t.Error("Expected false, go true")
	}
	if BinarySearch(haystack, 69421) {
		t.Error("Expected false, go true")
	}
	if BinarySearch(haystack, 2) {
		t.Error("Expected false, go true")
	}

}

func TestTwoCrystalBalls(t *testing.T) {
	arrayLen := 10_000
	idx := rand.Intn(arrayLen)

	breaks := make([]bool, arrayLen)
	for i := idx; i < arrayLen; i++ {
		breaks[i] = true
	}
	res := TwoCrystalBalls(breaks)
	if res != idx {
		t.Errorf("Expected %d, got %d", idx, res)
	}

	res = TwoCrystalBalls(make([]bool, 800))
	if res != -1 {
		t.Errorf("Expected -1, got %d", res)
	}
}
