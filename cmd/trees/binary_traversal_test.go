package trees

import (
	"reflect"
	"testing"
)

func TestBinaryPreOrder(t *testing.T) {
	path := BinaryPreOrderTraverse(ExampleBinaryTree)
	expected := []any{
		20,
		10,
		5,
		7,
		15,
		50,
		30,
		29,
		45,
		100,
	}

	if !reflect.DeepEqual(path, expected) {
		t.Error("Expected:", expected, "Received:", path)
	}
}

func TestBinaryInOrder(t *testing.T) {
	path := BinaryInOrderTraverse(ExampleBinaryTree)
	expected := []any{
		5,
		7,
		10,
		15,
		20,
		29,
		30,
		45,
		50,
		100,
	}

	if !reflect.DeepEqual(path, expected) {
		t.Error("Expected:", expected, "Received:", path)
	}
}

func TestBinaryPostOrder(t *testing.T) {
	path := BinaryPostOrderTraverse(ExampleBinaryTree)
	expected := []any{
		7,
		5,
		15,
		10,
		29,
		45,
		30,
		100,
		50,
		20,
	}

	if !reflect.DeepEqual(path, expected) {
		t.Error("Expected:", expected, "Received:", path)
	}
}
