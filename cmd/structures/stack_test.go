package structures_test

import (
	"testing"

	"skabillium.io/kata-go/cmd/structures"
)

func TestStack(t *testing.T) {
	s := structures.NewStack()

	s.Push(5)
	s.Push(7)
	s.Push(9)

	if s.Pop() != 9 {
		t.Error("Expected Pop() to return 9")
	}
	if s.Length != 2 {
		t.Error("Expected Length to be 2")
	}

	s.Push(11)
	if s.Pop() != 11 {
		t.Error("Expected Pop() to return 11")
	}
	if s.Pop() != 7 {
		t.Error("Expected Pop() to return 7")
	}
	if s.Peek() != 5 {
		t.Error("Expected Peek() to return 5")
	}
	if s.Pop() != 5 {
		t.Error("Expected Pop() to return 5")
	}
	if s.Pop() != nil {
		t.Error("Expected Pop() to return nil")
	}

	// Test after emptying
	s.Push(42)
	if s.Peek() != 42 {
		t.Error("Expected Peek() to return 42")
	}
	if s.Length != 1 {
		t.Error("Expected Length to be 1")
	}
}
