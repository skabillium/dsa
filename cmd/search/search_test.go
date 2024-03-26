package search_test

import (
	"math/rand"
	"reflect"
	"testing"

	"skabillium.io/kata-go/cmd/graphs"
	"skabillium.io/kata-go/cmd/search"
	"skabillium.io/kata-go/cmd/trees"
)

func TestLinearSearch(t *testing.T) {
	haystack := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	if !search.LinearSearch(haystack, 69) {
		t.Error("Expected true, go false")
	}
	if !search.LinearSearch(haystack, 69420) {
		t.Error("Expected true, go false")
	}
	if !search.LinearSearch(haystack, 1) {
		t.Error("Expected true, go false")
	}

	if search.LinearSearch(haystack, 1336) {
		t.Error("Expected false, go true")
	}
	if search.LinearSearch(haystack, 69421) {
		t.Error("Expected false, go true")
	}
	if search.LinearSearch(haystack, 2) {
		t.Error("Expected false, go true")
	}
}

func TestBinarySearch(t *testing.T) {
	haystack := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	if !search.BinarySearch(haystack, 69) {
		t.Error("Expected true, go false")
	}
	if !search.BinarySearch(haystack, 69420) {
		t.Error("Expected true, go false")
	}
	if !search.BinarySearch(haystack, 1) {
		t.Error("Expected true, go false")
	}

	if search.BinarySearch(haystack, 1336) {
		t.Error("Expected false, go true")
	}
	if search.BinarySearch(haystack, 69421) {
		t.Error("Expected false, go true")
	}
	if search.BinarySearch(haystack, 2) {
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
	res := search.TwoCrystalBalls(breaks)
	if res != idx {
		t.Errorf("Expected %d, got %d", idx, res)
	}

	res = search.TwoCrystalBalls(make([]bool, 800))
	if res != -1 {
		t.Errorf("Expected -1, got %d", res)
	}
}

func TestBinaryBfs(t *testing.T) {
	if !search.BinaryBfs(*trees.ExampleBinaryTree, 45) {
		t.Error("Expected BinaryBfs(45) to return true")
	}
	if !search.BinaryBfs(*trees.ExampleBinaryTree, 7) {
		t.Error("Expected BinaryBfs(7) to return true")
	}
	if search.BinaryBfs(*trees.ExampleBinaryTree, 70) {
		t.Error("Expected BinaryBfs(70) to return false")
	}
}

func TestBinarySearchTreeDfs(t *testing.T) {
	if !search.BinarySearchTreeDfs(trees.ExampleBinaryTree, 45) {
		t.Error("Expected BinaryBfs(45) to return true")
	}
	if !search.BinarySearchTreeDfs(trees.ExampleBinaryTree, 7) {
		t.Error("Expected BinaryBfs(7) to return true")
	}
	if search.BinarySearchTreeDfs(trees.ExampleBinaryTree, 70) {
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
	res := search.BfsGraphMatrix(graphs.Matrix2, 0, 6)

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
	res := search.DfsGraphList(graphs.List2, 0, 6)
	if !reflect.DeepEqual(res, expected) {
		t.Error("Expected DfsGraphList(graphs.List2, 0, 6) to return", expected, "got", res)
	}

	res = search.DfsGraphList(graphs.List2, 6, 0)
	if !reflect.DeepEqual(res, []int{}) {
		t.Error("Expected DfsGraphList(graphs.List2, 6, 0) to return", []int{}, "got", res)
	}
}
