package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{42, 1, 3, 8, 69, 10, 20}
	sorted := []int{1, 3, 8, 10, 20, 42, 69}
	BubbleSort(arr)
	if !reflect.DeepEqual(arr, sorted) {
		t.Error("Expected", sorted, "got", arr)
	}
}

func TestQuickSort(t *testing.T) {
	arr := []int{42, 1, 3, 8, 69, 10, 20}
	sorted := []int{1, 3, 8, 10, 20, 42, 69}
	QuickSort(arr)

	if !reflect.DeepEqual(arr, sorted) {
		t.Error("Expected", sorted, "got", arr)
	}
}
