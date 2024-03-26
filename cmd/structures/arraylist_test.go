package structures_test

import (
	"testing"

	"skabillium.io/kata-go/cmd/structures"
)

func TestArrayList(t *testing.T) {
	list := structures.NewArrayList(16)

	list.Append(5)
	list.Append(7)
	list.Append(9)

	if list.Get(2) != 9 {
		t.Error("Expected Get(2) to return 9")
	}
	if list.RemoveAt(1) != 7 {
		t.Error("Expected RemoveAt(1) to return 7")
	}
	if list.Length != 2 {
		t.Error("Expected Length to be 2")
	}

	list.Append(11)
	if list.RemoveAt(1) != 9 {
		t.Error("Expected RemoveAt(1) to return 9")
	}
	if list.Remove(9) != nil {
		t.Error("Expected Remove(9) to return nil")
	}
	if list.RemoveAt(0) != 5 {
		t.Error("Expected RemoveAt(0) to return 5")
	}
	if list.RemoveAt(0) != 11 {
		t.Error("Expected RemoveAt(0) to return 11")
	}
	if list.Length != 0 {
		t.Error("Expected Length to be 0")
	}

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)

	if list.Get(2) != 5 {
		t.Error("Expected Get(2) to return 5")
	}
	if list.Get(0) != 9 {
		t.Error("Expected Get(0) to return 9")
	}
	if list.Remove(9) != 9 {
		t.Error("Expected Remove(9) to return 9")
	}
	if list.Length != 2 {
		t.Error("Expected Length to be 2")
	}
	if list.Get(0) != 7 {
		t.Error("Expected Get(0) to return 7")
	}
}
