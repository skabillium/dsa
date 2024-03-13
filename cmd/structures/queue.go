package structures

type QueueNode struct {
	value any
	next  *QueueNode
}

type Queue struct {
	Length int

	head *QueueNode
	tail *QueueNode
}

func NewQueue() *Queue {
	return &Queue{Length: 0, head: nil, tail: nil}
}

func (q *Queue) Enqueue(item any) {
	node := &QueueNode{value: item, next: nil}
	if q.Length == 0 {
		q.head, q.tail = node, node
	}

	q.Length++
	q.tail.next = node
	q.tail = node
}

func (q *Queue) Deque() any {
	if q.head == nil {
		return nil
	}

	head := q.head
	q.head = q.head.next
	q.Length--

	return head.value
}

func (q *Queue) Peek() any {
	if q.head == nil {
		return nil
	}
	return q.head.value
}
