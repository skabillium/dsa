package structures

import "testing"

func TestRingBuffer(t *testing.T) {
	buffer := NewRingBuffer(8)

	buffer.Enqueue(5)
	if buffer.Deque() != 5 {
		t.Errorf("Expected Deque() to return 5")
	}
	if buffer.Deque() != nil {
		t.Errorf("Expected Deque() to return nil")
	}

	buffer.Enqueue(42)
	buffer.Enqueue(9)
	if buffer.Deque() != 42 {
		t.Errorf("Expected Deque() to return 42")
	}
	if buffer.Deque() != 9 {
		t.Errorf("Expected Deque() to return 9")
	}
	if buffer.Deque() != nil {
		t.Errorf("Expected Deque() to return nil")
	}

	buffer.Enqueue(42)
	buffer.Enqueue(7)
	buffer.Enqueue(3)
	if buffer.Get(2) != 3 {
		t.Error("Expected Get(2) to return 3")
	}
	if buffer.Get(1) != 7 {
		t.Error("Expected Get(2) to return 7")
	}
	if buffer.Get(0) != 42 {
		t.Error("Expected Get(2) to return 42")
	}
}
