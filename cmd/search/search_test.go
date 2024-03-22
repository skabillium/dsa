package search

import (
	"math/rand"
	"reflect"
	"testing"

	"skabillium.io/kata-go/cmd/graphs"
	"skabillium.io/kata-go/cmd/trees"
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

func TestBinaryBfs(t *testing.T) {
	if !BinaryBfs(*trees.ExampleBinaryTree, 45) {
		t.Error("Expected BinaryBfs(45) to return true")
	}
	if !BinaryBfs(*trees.ExampleBinaryTree, 7) {
		t.Error("Expected BinaryBfs(7) to return true")
	}
	if BinaryBfs(*trees.ExampleBinaryTree, 70) {
		t.Error("Expected BinaryBfs(70) to return false")
	}
}

func TestBinarySearchTreeDfs(t *testing.T) {
	if !BinarySearchTreeDfs(trees.ExampleBinaryTree, 45) {
		t.Error("Expected BinaryBfs(45) to return true")
	}
	if !BinarySearchTreeDfs(trees.ExampleBinaryTree, 7) {
		t.Error("Expected BinaryBfs(7) to return true")
	}
	if BinarySearchTreeDfs(trees.ExampleBinaryTree, 70) {
		t.Error("Expected BinaryBfs(70) to return false")
	}
}

func TestBfsGraphMatrix(t *testing.T) {
	expected := []int{
		0,
		1,
		4,
		5,
		6,
	}
	res := BfsGraphMatrix(graphs.Matrix2, 0, 6)

	if !reflect.DeepEqual(res, expected) {
		t.Error("Expected BfsGraphMatrix(graphs.Matrix2, 0, 6) to return", expected, "got", res)
	}
}

func TestDfsGraphList(t *testing.T) {
	expected := []int{
		0,
		1,
		4,
		5,
		6,
	}
	res := DfsGraphList(graphs.List2, 0, 6)
	if !reflect.DeepEqual(res, expected) {
		t.Error("Expected DfsGraphList(graphs.List2, 0, 6) to return", expected, "got", res)
	}

	res = DfsGraphList(graphs.List2, 6, 0)
	if !reflect.DeepEqual(res, []int{}) {
		t.Error("Expected DfsGraphList(graphs.List2, 6, 0) to return", []int{}, "got", res)
	}
}
