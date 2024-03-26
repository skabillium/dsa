package structures_test

import (
	"testing"

	"skabillium.io/kata-go/cmd/structures"
)

func TestQueue(t *testing.T) {
	q := structures.NewQueue()

	q.Enqueue(5)
	q.Enqueue(7)
	q.Enqueue(9)

	if q.Deque() != 5 {
		t.Error("Expected Deque() to return 5")
	}
	if q.Length != 2 {
		t.Error("Expected Length to be 2")
	}

	q.Enqueue(11)

	if q.Deque() != 7 {
		t.Error("Expected Deque() to return 7")
	}
	if q.Deque() != 9 {
		t.Error("Expected Deque() to return 9")
	}
	if q.Peek() != 11 {
		t.Error("Expected Peek() to return 11")
	}
	if q.Deque() != 11 {
		t.Error("Expected Deque() to return 11")
	}
	if q.Deque() != nil {
		t.Error("Expected Deque() to return nil")
	}
	if q.Length != 0 {
		t.Error("Expected Length to be 0")
	}

	// Add an element after emptying
	q.Enqueue(42)
	if q.Peek() != 42 {
		t.Error("Expected Peek() to return 42")
	}
	if q.Length != 1 {
		t.Error("Expected Length to be 1")
	}
}
