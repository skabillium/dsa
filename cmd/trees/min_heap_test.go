package trees

import "testing"

func TestMinHeap(t *testing.T) {
	// const heap = new MinHeap();
	heap := NewMinHeap()

	// expect(heap.length).toEqual(0);
	if heap.Length != 0 {
		t.Error("Expected Length to be 0")
	}

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)

	// expect(heap.length).toEqual(8);
	if heap.Length != 8 {
		t.Error("Expected Length to be 8")
	}
	// expect(heap.delete()).toEqual(1);
	if heap.Delete() != 1 {
		t.Error("Expected Delete() to return 1")
	}
	// expect(heap.delete()).toEqual(3);
	if heap.Delete() != 3 {
		t.Error("Expected Delete() to return 3")
	}
	// expect(heap.delete()).toEqual(4);
	if heap.Delete() != 4 {
		t.Error("Expected Delete() to return 4")
	}
	// expect(heap.delete()).toEqual(5);
	if heap.Delete() != 5 {
		t.Error("Expected Delete() to return 5")
	}
	// expect(heap.length).toEqual(4);
	if heap.Length != 4 {
		t.Error("Expected Length to be 4")
	}
	// expect(heap.delete()).toEqual(7);
	if heap.Delete() != 7 {
		t.Error("Expected Delete() to return 7")
	}
	// expect(heap.delete()).toEqual(8);
	if heap.Delete() != 8 {
		t.Error("Expected Delete() to return 8")
	}
	// expect(heap.delete()).toEqual(69);
	if heap.Delete() != 69 {
		t.Error("Expected Delete() to return 69")
	}
	// expect(heap.delete()).toEqual(420);
	if heap.Delete() != 420 {
		t.Error("Expected Delete() to return 420")
	}
	// expect(heap.length).toEqual(0);
	if heap.Length != 0 {
		t.Error("Expected Length to be 0")
	}
}
