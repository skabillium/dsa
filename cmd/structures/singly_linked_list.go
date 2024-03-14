package structures

import "fmt"

type SllNode struct {
	value any
	next  *SllNode
}

type SinglyLinkedList struct {
	Length int
	head   *SllNode
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{Length: 0, head: nil}
}

func (l *SinglyLinkedList) Append(item any) {
	if l.Length == 0 {
		l.Prepend(item)
		return
	}

	curr := l.head
	for i := 0; i < l.Length-1; i++ {
		curr = curr.next
	}

	l.Length++
	node := &SllNode{value: item}
	curr.next = node
}

func (l *SinglyLinkedList) Prepend(item any) {
	node := &SllNode{value: item}
	if l.Length == 0 {
		l.Length++
		l.head = node
		return
	}

	l.Length++
	node.next = l.head
	l.head = node
}

func (l *SinglyLinkedList) InsertAt(idx int, item any) {
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
	for i := 0; i < idx-1; i++ {
		curr = curr.next
	}

	l.Length++
	node := &SllNode{value: item}
	node.next = curr.next
	curr.next = node
}

func (l *SinglyLinkedList) Get(idx int) any {
	if l.Length == 0 || idx < 0 || idx > l.Length {
		return nil
	}

	curr := l.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}

	return curr.value
}

func (l *SinglyLinkedList) Remove(item any) any {
	if l.Length == 0 {
		return nil
	}

	if l.head.value == item {
		return l.RemoveAt(0)
	}

	curr := l.head
	for i := 0; i < l.Length-1; i++ {
		if curr.next.value == item {
			l.Length--
			value := curr.next.value
			curr.next = curr.next.next
			return value
		}
		curr = curr.next
	}

	return nil
}

func (l *SinglyLinkedList) RemoveAt(idx int) any {
	if idx < 0 || idx >= l.Length {
		return nil
	}

	if idx == 0 {
		l.Length--
		value := l.head.value
		l.head = l.head.next
		return value
	}

	curr := l.head
	for i := 0; i < idx-1; i++ {
		curr = curr.next
	}

	l.Length--
	value := curr.next.value
	curr.next = curr.next.next
	return value
}

func (l *SinglyLinkedList) Debug() {
	values := []any{}

	if l.Length == 0 {
		fmt.Println(values...)
		return
	}

	curr := l.head
	for i := 0; i < l.Length; i++ {
		values = append(values, curr.value)
		curr = curr.next
	}

	fmt.Println(values...)
}
