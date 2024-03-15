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
