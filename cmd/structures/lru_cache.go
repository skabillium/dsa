package structures

type LruNode struct {
	Value any
	Key   any
	Next  *LruNode
	Prev  *LruNode
}

type LruCache struct {
	length   int
	capacity int
	head     *LruNode
	tail     *LruNode
	lookup   map[any]*LruNode
}

func NewLruCache(capacity int) *LruCache {
	return &LruCache{
		length:   0,
		capacity: capacity,
		head:     nil,
		tail:     nil,
		lookup:   make(map[any]*LruNode),
	}
}

func (l *LruCache) Update(key any, value any) {
	node, found := l.lookup[key]
	if !found {
		node = &LruNode{
			Key:   key,
			Value: value,
		}
		l.prepend(node)
		l.length++
		l.trim()
		l.lookup[key] = node
		return
	}

	l.detach(node)
	l.prepend(node)
	node.Value = value
}

func (l *LruCache) Get(key any) any {
	node, found := l.lookup[key]
	if !found {
		return nil
	}

	l.detach(node)
	l.prepend(node)
	return node.Value
}

func (l *LruCache) prepend(node *LruNode) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.Next = l.head
	l.head.Prev = node
}

func (l *LruCache) detach(node *LruNode) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if l.head == node {
		l.head = nil
	}

	if l.tail == node {
		l.tail = nil
	}

	node.Next = nil
	node.Prev = nil

	// TODO: Maybe update length
}

func (l *LruCache) trim() {
	if l.length <= l.capacity {
		return
	}

	tail := l.tail
	l.detach(tail)
	delete(l.lookup, tail.Key)
}
