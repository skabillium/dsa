package structures

type RingBuffer struct {
	Capacity int
	arr      []any
	head     int
	tail     int
}

func NewRingBuffer(capacity int) *RingBuffer {
	return &RingBuffer{
		Capacity: capacity,
		arr:      make([]any, capacity),
		head:     0,
		tail:     0,
	}
}

// func (rb *RingBuffer) grow() {
// 	capacity := 2 * rb.Capacity
// 	arr := make([]any, capacity)

// 	copy(arr[:rb.Capacity], rb.arr[:])
// 	rb.arr = arr
// }

func (rb *RingBuffer) increment(counter int) int {
	return (counter + 1) % rb.Capacity
}

func (rb *RingBuffer) Enqueue(item any) {
	rb.arr[rb.tail] = item
	rb.tail = rb.increment(rb.tail)
}

func (rb *RingBuffer) Deque() any {
	if rb.head == rb.tail {
		return nil
	}

	item := rb.arr[rb.head]
	rb.arr[rb.head] = nil
	rb.head = rb.increment(rb.head)

	return item
}

func (rb *RingBuffer) Peek() any {
	return rb.arr[rb.head]
}

func (rb *RingBuffer) Get(idx int) any {
	pos := rb.head + idx
	if pos < 0 || pos > rb.Capacity-1 {
		return nil
	}

	return rb.arr[pos]
}
