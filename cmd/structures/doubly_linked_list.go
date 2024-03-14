package structures

type DllNode struct {
	value any
	prev  *DllNode
	next  *DllNode
}

type DoublyLinkedList struct {
	Length int
	head   *DllNode
	tail   *DllNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{Length: 0, head: nil, tail: nil}
}

func (l *DoublyLinkedList) Prepend(item any) {
	node := &DllNode{value: item}

	l.Length++
	if l.Length == 1 {
		l.head = node
		l.tail = node
		return
	}

	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *DoublyLinkedList) Append(item any) {
	l.Length++
	node := &DllNode{value: item}
	if l.Length == 1 {
		l.head = node
		l.tail = node
		return
	}

	l.tail.next = node
	node.prev = l.tail
	l.tail = node
}

func (l *DoublyLinkedList) InsertAt(idx int, item any) {
	if idx < 0 || idx > l.Length {
		return
	}

	if idx == 0 {
		l.Prepend(item)
		return
	}

	if idx == l.Length {
		l.Append(item)
		return
	}

	curr := l.head
	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.next
	}

	l.Length++
	node := &DllNode{value: item}
	node.prev = curr.prev
	node.next = curr

	curr.prev.next = node
	curr.prev = node
}

func (l *DoublyLinkedList) Get(idx int) any {
	node := l.getAt(idx)
	if node == nil {
		return nil
	}

	return node.value
}

func (l *DoublyLinkedList) Remove(item any) any {
	curr := l.head
	for i := 0; curr != nil && i < l.Length; i++ {
		if curr.value == item {
			break
		}
		curr = curr.next
	}

	if curr == nil {
		return nil
	}

	return l.removeNode(curr)
}

func (l *DoublyLinkedList) RemoveAt(idx int) any {
	node := l.getAt(idx)
	if node == nil {
		return nil
	}

	l.removeNode(node)
	return node.value
}

func (l *DoublyLinkedList) getAt(idx int) *DllNode {
	curr := l.head
	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.next
	}

	return curr
}

func (l *DoublyLinkedList) removeNode(node *DllNode) any {
	l.Length--
	if l.Length == 0 {
		value := l.head.value
		l.head = nil
		l.tail = nil
		return value
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == l.head {
		l.head = node.next
	}

	if node == l.tail {
		l.tail = node.prev
	}

	node.next = nil
	node.prev = nil
	return node.value
}
